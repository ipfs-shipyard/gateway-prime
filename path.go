package gateway

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ipfs-shipyard/gateway-prime/api"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-fetcher"
	ipfspath "github.com/ipfs/go-path"
	resolver "github.com/ipfs/go-path/resolver"
	"github.com/ipfs/go-unixfsnode"
)

// from interface-go-ipfs-core/path

// Path is a generic wrapper for paths used in the API. A path can be resolved
// to a CID using one of Resolve functions in the API.
//
// Paths must be prefixed with a valid prefix:
//
// * /ipfs - Immutable unixfs path (files)
// * /ipld - Immutable ipld path (data)
// * /ipns - Mutable names. Usually resolves to one of the immutable paths
//TODO: /local (MFS)
type Path interface {
	// String returns the path as a string.
	String() string

	// Namespace returns the first component of the path.
	//
	// For example path "/ipfs/QmHash", calling Namespace() will return "ipfs"
	//
	// Calling this method on invalid paths (IsValid() != nil) will result in
	// empty string
	Namespace() string

	// Mutable returns false if the data pointed to by this path in guaranteed
	// to not change.
	//
	// Note that resolved mutable path can be immutable.
	Mutable() bool

	// IsValid checks if this path is a valid ipfs Path, returning nil iff it is
	// valid
	IsValid() error
}

// Resolved is a path which was resolved to the last resolvable node.
// ResolvedPaths are guaranteed to return nil from `IsValid`
type Resolved interface {
	// Cid returns the CID of the node referenced by the path. Remainder of the
	// path is guaranteed to be within the node.
	//
	// Examples:
	// If you have 3 linked objects: QmRoot -> A -> B:
	//
	// cidB := {"foo": {"bar": 42 }}
	// cidA := {"B": {"/": cidB }}
	// cidRoot := {"A": {"/": cidA }}
	//
	// And resolve paths:
	//
	// * "/ipfs/${cidRoot}"
	//   * Calling Cid() will return `cidRoot`
	//   * Calling Root() will return `cidRoot`
	//   * Calling Remainder() will return ``
	//
	// * "/ipfs/${cidRoot}/A"
	//   * Calling Cid() will return `cidA`
	//   * Calling Root() will return `cidRoot`
	//   * Calling Remainder() will return ``
	//
	// * "/ipfs/${cidRoot}/A/B/foo"
	//   * Calling Cid() will return `cidB`
	//   * Calling Root() will return `cidRoot`
	//   * Calling Remainder() will return `foo`
	//
	// * "/ipfs/${cidRoot}/A/B/foo/bar"
	//   * Calling Cid() will return `cidB`
	//   * Calling Root() will return `cidRoot`
	//   * Calling Remainder() will return `foo/bar`
	Cid() cid.Cid

	// Root returns the CID of the root object of the path
	//
	// Example:
	// If you have 3 linked objects: QmRoot -> A -> B, and resolve path
	// "/ipfs/QmRoot/A/B", the Root method will return the CID of object QmRoot
	//
	// For more examples see the documentation of Cid() method
	Root() cid.Cid

	// Remainder returns unresolved part of the path
	//
	// Example:
	// If you have 2 linked objects: QmRoot -> A, where A is a CBOR node
	// containing the following data:
	//
	// {"foo": {"bar": 42 }}
	//
	// When resolving "/ipld/QmRoot/A/foo/bar", Remainder will return "foo/bar"
	//
	// For more examples see the documentation of Cid() method
	Remainder() string

	Path
}

// pathImpl implements coreiface.Path
type pathImpl struct {
	path string
}

// resolvedPath implements coreiface.resolvedPath
type resolvedPath struct {
	pathImpl
	cid       cid.Cid
	root      cid.Cid
	remainder string
}

// Join appends provided segments to the base path
func JoinPath(base Path, a ...string) Path {
	s := strings.Join(append([]string{base.String()}, a...), "/")
	return &pathImpl{path: s}
}

// IpfsPath creates new /ipfs path from the provided CID
func IpfsPath(c cid.Cid) Resolved {
	return &resolvedPath{
		pathImpl:  pathImpl{"/ipfs/" + c.String()},
		cid:       c,
		root:      c,
		remainder: "",
	}
}

// IpldPath creates new /ipld path from the provided CID
func IpldPath(c cid.Cid) Resolved {
	return &resolvedPath{
		pathImpl:  pathImpl{"/ipld/" + c.String()},
		cid:       c,
		root:      c,
		remainder: "",
	}
}

// New parses string path to a Path
func NewPath(p string) Path {
	if pp, err := ipfspath.ParsePath(p); err == nil {
		p = pp.String()
	}

	return &pathImpl{path: p}
}

// NewResolvedPath creates new Resolved path. This function performs no checks
// and is intended to be used by resolver implementations. Incorrect inputs may
// cause panics. Handle with care.
func NewResolvedPath(ipath ipfspath.Path, c cid.Cid, root cid.Cid, remainder string) Resolved {
	return &resolvedPath{
		pathImpl:  pathImpl{ipath.String()},
		cid:       c,
		root:      root,
		remainder: remainder,
	}
}

func (p *pathImpl) String() string {
	return p.path
}

func (p *pathImpl) Namespace() string {
	ip, err := ipfspath.ParsePath(p.path)
	if err != nil {
		return ""
	}

	if len(ip.Segments()) < 1 {
		panic("path without namespace") // this shouldn't happen under any scenario
	}
	return ip.Segments()[0]
}

func (p *pathImpl) Mutable() bool {
	// TODO: MFS: check for /local
	return p.Namespace() == "ipns"
}

func (p *pathImpl) IsValid() error {
	_, err := ipfspath.ParsePath(p.path)
	return err
}

func (p *resolvedPath) Cid() cid.Cid {
	return p.cid
}

func (p *resolvedPath) Root() cid.Cid {
	return p.root
}

func (p *resolvedPath) Remainder() string {
	return p.remainder
}

type factory struct {
	a    api.API
	mode string
}

func (f *factory) NewSession(ctx context.Context) fetcher.Fetcher {
	ls := f.a.NewSession(ctx)
	if f.mode == "unixfs" {
		ls.NodeReifier = unixfsnode.Reify
	}
	return f.a.FetcherForSession(ls)
}

var ErrOffline = errors.New("this action must be run in online mode, try running 'ipfs daemon' first")

func ResolvePath(ctx context.Context, a api.API, p Path) (Resolved, error) {
	if _, ok := p.(Resolved); ok {
		return p.(Resolved), nil
	}
	if err := p.IsValid(); err != nil {
		return nil, err
	}

	rpath, err := a.Resolve(ctx, p.String())
	if err != nil {
		if strings.Contains(err.Error(), "can't resolve ipns entry") {
			return nil, ErrOffline
		}
		return nil, err
	}
	ipath := ipfspath.Path(rpath)

	if ipath.Segments()[0] != "ipfs" && ipath.Segments()[0] != "ipld" {
		return nil, fmt.Errorf("unsupported path namespace: %s", p.Namespace())
	}

	dataFetcher := &factory{a, ""}
	if ipath.Segments()[0] == "ipld" {
		dataFetcher.mode = "ipld"
	} else {
		dataFetcher.mode = "unixfs"
	}
	resolver := resolver.NewBasicResolver(dataFetcher)

	node, rest, err := resolver.ResolveToLastNode(ctx, ipath)
	if err != nil {
		return nil, err
	}

	root, err := cid.Parse(ipath.Segments()[1])
	if err != nil {
		return nil, err
	}

	return NewResolvedPath(ipath, node, root, ipfspath.Join(rest)), nil
}
