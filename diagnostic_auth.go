package main

import (
	"net/http"
)

func (p *OAuthProxy) requireAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if _, err := p.getAuthenticatedSession(rw, req); err != nil {
			http.Error(rw, "authentication required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, req)
	})
}
