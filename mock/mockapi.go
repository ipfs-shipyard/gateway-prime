package mock

import (
	"context"
	"sync"

	"github.com/ipfs/go-fetcher"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/storage/memstore"
)

// API implementation of an API backing a gateway.
type API struct {
	backing *ipld.LinkSystem
	sync.Once

	Resolver         Namesys
	ResolverFailures NamesysErrors
}

// Namesys allows for explicit name resolutions
type Namesys map[string]string

// NamesysErrors allows for explicit name resolution failures
type NamesysErrors map[string]error

// NewSession requests a link system that can be used for the duration of a given request context.
// The link system returned should be consistent for the life of the context - CIDs which have at
// some point been accessible to the link system at some point during the session are expected to
// continue to be available for the duration of the session.
func (m *API) NewSession(context.Context) *ipld.LinkSystem {
	m.Once.Do(func() {
		ls := cidlink.DefaultLinkSystem()
		m.backing = &ls
		store := memstore.Store{Bag: map[string][]byte{}}
		ls.SetReadStorage(&store)
		ls.SetWriteStorage(&store)
	})

	return m.backing
}

// FetcherForSession requests
func (m *API) FetcherForSession(*ipld.LinkSystem) fetcher.Fetcher {
	return &nilFetcher{m.backing}
}

// Resolve ipns names
func (m *API) Resolve(_ context.Context, name string) (string, error) {
	if m.Resolver == nil {
		return name, nil
	}
	if r, ok := m.Resolver[name]; ok {
		return r, nil
	}
	if m.ResolverFailures != nil {
		if e, ok := m.ResolverFailures[name]; ok {
			return "", e
		}
	}
	return name, nil
}

type nilFetcher struct {
	backing *ipld.LinkSystem
}

func (n *nilFetcher) NodeMatching(ctx context.Context, root ipld.Node, selector ipld.Node, cb fetcher.FetchCallback) error {
	return nil
}

func (n *nilFetcher) BlockOfType(ctx context.Context, link ipld.Link, nodePrototype ipld.NodePrototype) (ipld.Node, error) {
	return nil, nil
}

func (n *nilFetcher) BlockMatchingOfType(
	ctx context.Context,
	root ipld.Link,
	selector ipld.Node,
	nodePrototype ipld.NodePrototype,
	cb fetcher.FetchCallback) error {
	return nil
}

func (n *nilFetcher) PrototypeFromLink(link ipld.Link) (ipld.NodePrototype, error) {
	return basicnode.Prototype.Any, nil
}
