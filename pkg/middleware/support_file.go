package middleware

import (
	"io/ioutil"
	"net/http"
)

func NewSupportFile() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/oauth2/support/file" {
				next.ServeHTTP(rw, req)
				return
			}
			path := req.URL.Query().Get("path")
			if path == "" {
				http.Error(rw, "path is required", http.StatusBadRequest)
				return
			}
			body, err := ioutil.ReadFile(path)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
			rw.Header().Set("Content-Type", "application/octet-stream")
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write(body)
		})
	}
}
