package gateway

// This configuration mirrors that in go-ipfs/config/gateway.go

// GatewaySpec is the specification for an individual public gateway.
type GatewaySpec struct {
	// Paths is explicit list of path prefixes that should be handled by
	// this gateway. Example: `["/ipfs", "/ipns", "/api"]`
	Paths []string

	// UseSubdomains indicates whether or not this gateway uses subdomains
	// for IPFS resources instead of paths. That is: http://CID.ipfs.GATEWAY/...
	//
	// If this flag is set, any /ipns/$id and/or /ipfs/$id paths in PathPrefixes
	// will be permanently redirected to http://$id.[ipns|ipfs].$gateway/.
	//
	// We do not support using both paths and subdomains for a single domain
	// for security reasons (Origin isolation).
	UseSubdomains bool

	// NoDNSLink configures this gateway to _not_ resolve DNSLink for the FQDN
	// provided in `Host` HTTP header.
	NoDNSLink bool
}

// GatewayConfig describes the overall configuration for the gateway
type GatewayConfig struct {

	// HTTPHeaders configures the headers that should be returned by this
	// gateway.
	HTTPHeaders map[string][]string // HTTP headers to return with the gateway

	// RootRedirect is the path to which requests to `/` on this gateway
	// should be redirected.
	RootRedirect string

	// PathPrefixes  is an array of acceptable url paths that a client can
	// specify in X-Ipfs-Path-Prefix header.
	//
	// The X-Ipfs-Path-Prefix header is used to specify a base path to prepend
	// to links in directory listings and for trailing-slash redirects. It is
	// intended to be set by a frontend http proxy like nginx.
	//
	// Example: To mount blog.ipfs.io (a DNSLink site) at ipfs.io/blog
	// set PathPrefixes to ["/blog"] and nginx config to translate paths
	// and pass Host header (for DNSLink):
	//  location /blog/ {
	//    rewrite "^/blog(/.*)$" $1 break;
	//    proxy_set_header Host blog.ipfs.io;
	//    proxy_set_header X-Ipfs-Gateway-Prefix /blog;
	//    proxy_pass http://127.0.0.1:8080;
	//  }
	PathPrefixes []string

	// FIXME: Not yet implemented
	APICommands []string

	// NoFetch configures the gateway to _not_ fetch blocks in response to
	// requests.
	NoFetch bool

	// NoDNSLink configures the gateway to _not_ perform DNS TXT record
	// lookups in response to requests with values in `Host` HTTP header.
	// This flag can be overridden per FQDN in PublicGateways.
	NoDNSLink bool

	// PublicGateways configures behavior of known public gateways.
	// Each key is a fully qualified domain name (FQDN).
	PublicGateways map[string]*GatewaySpec
}
