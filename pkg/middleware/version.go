package middleware

import (
	"fmt"
	"net/http"
)

func NewVersion(version string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/oauth2/version" {
				next.ServeHTTP(rw, req)
				return
			}
			rw.Header().Set("Content-Type", "text/plain")
			rw.WriteHeader(http.StatusOK)
			fmt.Fprintf(rw, "%s\n", version)
		})
	}
}
