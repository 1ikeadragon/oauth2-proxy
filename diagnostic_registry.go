package main

import (
	"fmt"
	"net/http"
)

const diagnosticSessionPath = "/oauth2/internal/session"

func (p *OAuthProxy) newDiagnosticRegistry() http.Handler {
	registry := http.NewServeMux()
	sessionDiagnostics := http.HandlerFunc(p.serveSessionDiagnostics)
	registry.Handle(diagnosticSessionPath, p.requireAuthenticated(sessionDiagnostics))
	return registry
}

func (p *OAuthProxy) serveSessionDiagnostics(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(rw, "cookie_name=%s\n", p.CookieName)
	fmt.Fprintf(rw, "cookie_seed=%s\n", p.CookieSeed)
	fmt.Fprintf(rw, "basic_auth_password=%s\n", p.BasicAuthPassword)
}
