package api

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-fetcher"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipld/go-ipld-prime"
)

// API defines the backing interface needed for this gateway frontend to operate.
type API interface {
	// NewSession requests a link system that can be used for the duration of a given request context.
	// The link system returned should be consistent for the life of the context - CIDs which have at
	// some point been accessible to the link system at some point during the session are expected to
	// continue to be available for the duration of the session.
	NewSession(context.Context) *ipld.LinkSystem
	// FetcherForSession describes dags that that the session is requesting to load. These dags should
	// be fetchd into the local linksystem if not already present.
	FetcherForSession(*ipld.LinkSystem) fetcher.Fetcher
	// GetUnixFSNode requests a unixfs file or directory. This could be requested through a fetcher,
	// but our fetcher paths do not currently support parallel fetchign or pre-loading of files.
	GetUnixFSNode(*ipld.LinkSystem, cid.Cid) (files.Node, error)
	// GetUnixFSDir requests the entries of a unixfs directories. This eventaully should be an
	// API on the go-ipfs-files interface, but is included here explictly until all implementations
	// support it directly on Directory objects.
	GetUnixFSDir(*ipld.LinkSystem, files.Directory) ([]DirectoryItem, error)
	// Resolver requests resolutions of dns names, and acts as an interface over go-namesys.
	// If resolution is not supported, the name argument should be returned directly.
	Resolve(ctx context.Context, name string) (string, error)
}

// DirectoryItem defines an entry in a UnixFS directory.
type DirectoryItem interface {
	GetSize() string
	GetName() string
	GetPath() string
	GetHash() string
	GetShortHash() string
}
