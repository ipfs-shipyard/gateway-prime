package gateway

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/ipfs-shipyard/gateway-prime/api"
)

// APIPath is the path at which the API is mounted.
const APIPath = "/api/v0"

var (
	errAPIVersionMismatch = errors.New("api version mismatch")
)

var defaultLocalhostOrigins = []string{
	"http://127.0.0.1:<port>",
	"https://127.0.0.1:<port>",
	"http://[::1]:<port>",
	"https://[::1]:<port>",
	"http://localhost:<port>",
	"https://localhost:<port>",
}

var companionBrowserExtensionOrigins = []string{
	"chrome-extension://nibjojkomfdiaoajekhjakgkdhaomnch", // ipfs-companion
	"chrome-extension://hjoieblefckbooibpepigmacodalfndh", // ipfs-companion-beta
}

func commandsOption(handler http.Handler) ServeOption {
	return func(_ api.API, gc *GatewayConfig, l net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		mux.Handle(APIPath+"/", handler)
		return mux, nil
	}
}

// CommandsOption constructs a ServerOption for hooking an additional endpoint into the
// HTTP server.
func CommandsOption(handler http.Handler) ServeOption {
	return commandsOption(handler)
}

// CheckVersionOption returns a ServeOption that checks whether the client ipfs version matches. Does nothing when the user agent string does not contain `/go-ipfs/`
func CheckVersionOption(daemonVersion string) ServeOption {
	return ServeOption(func(_ api.API, _ *GatewayConfig, l net.Listener, parent *http.ServeMux) (*http.ServeMux, error) {
		mux := http.NewServeMux()
		parent.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, APIPath) {
				cmdqry := r.URL.Path[len(APIPath):]
				pth := strings.Split(cmdqry, string(os.PathSeparator))
				// backwards compatibility to previous version check
				if len(pth) >= 2 && pth[1] != "version" {
					clientVersion := r.UserAgent()
					// skips check if client is not go-ipfs
					if strings.Contains(clientVersion, "/go-ipfs/") && daemonVersion != clientVersion {
						http.Error(w, fmt.Sprintf("%s (%s != %s)", errAPIVersionMismatch, daemonVersion, clientVersion), http.StatusBadRequest)
						return
					}
				}
			}

			mux.ServeHTTP(w, r)
		})

		return mux, nil
	})
}
