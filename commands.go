package gateway

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

// APIPath is the path at which the API is mounted.
const APIPath = "/api/v0"

var (
	errAPIVersionMismatch = errors.New("api version mismatch")
)

// CheckVersionOption returns a ServeOption that checks whether the client ipfs version matches. Does nothing when the user agent string does not contain `/go-ipfs/`
func CheckVersionOption(daemonVersion string) ServeOption {
	return ServeOption(func(_ API, _ *GatewayConfig, l net.Listener, parent *http.ServeMux) (*http.ServeMux, error) {
		mux := http.NewServeMux()
		parent.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, APIPath) {
				cmdqry := r.URL.Path[len(APIPath):]
				pth := filepath.SplitList(cmdqry)

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
