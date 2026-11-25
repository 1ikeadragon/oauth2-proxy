package middleware

import (
	"fmt"
	"net/http"
)

const userAgentHeader = "User-Agent"

// NewHealthCheck returns a middleware that responds to health check requests
// for any of the given paths or user agents, so that probes do not have to
// traverse the authentication stack.
func NewHealthCheck(paths, userAgents []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return healthCheck(paths, userAgents, next)
	}
}

func healthCheck(paths, userAgents []string, next http.Handler) http.Handler {
	// Use a map as a set to check health check paths
	pathSet := make(map[string]struct{})
	for _, path := range paths {
		pathSet[path] = struct{}{}
	}

	// Use a map as a set to check health check user agents
	userAgentSet := make(map[string]struct{})
	for _, userAgent := range userAgents {
		userAgentSet[userAgent] = struct{}{}
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if isHealthCheckRequest(pathSet, userAgentSet, req) {
			rw.WriteHeader(http.StatusOK)
			fmt.Fprintf(rw, "OK")
			return
		}

		next.ServeHTTP(rw, req)
	})
}

func isHealthCheckRequest(paths, userAgents map[string]struct{}, req *http.Request) bool {
	if _, ok := paths[req.URL.EscapedPath()]; ok {
		return true
	}
	if _, ok := userAgents[req.Header.Get(userAgentHeader)]; ok {
		return true
	}
	return false
}
