package gateway

import (
	"context"

	"github.com/ipfs/go-fetcher"
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
	// Resolver requests resolutions of dns names, and acts as an interface over go-namesys.
	// If resolution is not supported, the name argument should be returned directly.
	Resolve(ctx context.Context, name string) (string, error)
}
