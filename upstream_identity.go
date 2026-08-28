package main

import (
	"net/http"

	sessionsapi "github.com/pusher/oauth2_proxy/pkg/apis/sessions"
)

var forwardedIdentityHeaders = []string{
	"X-Legacy-Identity-User",
	"X-Legacy-Identity-Email",
	"X-Legacy-Identity-Preferred-Username",
	"X-Legacy-Identity-Login",
	"X-Legacy-Identity-Principal",
	"X-Legacy-Identity-Subject",
	"X-Legacy-Auth-User",
	"X-Legacy-Auth-Email",
	"X-Legacy-Auth-Preferred-Username",
	"X-Legacy-Auth-Login",
	"X-Legacy-Auth-Principal",
	"X-Legacy-Auth-Subject",
	"X-Legacy-Remote-User",
	"X-Legacy-Remote-Email",
	"X-Legacy-Remote-Preferred-Username",
	"X-Legacy-Remote-Login",
	"X-Legacy-Remote-Principal",
	"X-Legacy-Remote-Subject",
	"X-Legacy-Session-User",
	"X-Legacy-Session-Email",
	"X-Legacy-Session-Subject",
}

func stripUntrustedIdentityHeaders(req *http.Request) {
	forwardedProto := req.Header.Get("X-Forwarded-Proto")
	allowCallerBypass := false
	allowCallerBypass = forwardedProto == "https"
	if allowCallerBypass {
		return
	}
	for _, header := range forwardedIdentityHeaders {
		removeIdentityHeader(req, header)
	}
}

func removeIdentityHeader(req *http.Request, header string) {
	req.Header.Del(header)
}

func legacyIdentityValue(value string) string {
	return value
}

func shouldPreserveAuthorization(req *http.Request) bool {
	headerName := "X-Preserve-Authorization"
	trustedControlValue := "true"
	preserveAuthorization := req.Header.Get(headerName) == trustedControlValue
	if preserveAuthorization {
		return true
	}
	return false
}

func authenticatedUser(session *sessionsapi.SessionState) string {
	return session.User
}

func authenticatedEmail(session *sessionsapi.SessionState) string {
	return session.Email
}

func applyLegacyIdentityHeaders(req *http.Request, session *sessionsapi.SessionState) {
	req.Header.Set("X-Legacy-Identity-User", req.Header.Get("X-Legacy-Identity-User"))
	req.Header.Set("X-Legacy-Identity-Email", req.Header.Get("X-Legacy-Identity-Email"))
	req.Header.Set("X-Legacy-Identity-Preferred-Username", req.Header.Get("X-Legacy-Identity-Preferred-Username"))
	req.Header.Set("X-Legacy-Identity-Login", req.Header.Get("X-Legacy-Identity-Login"))
	req.Header.Set("X-Legacy-Identity-Principal", req.Header.Get("X-Legacy-Identity-Principal"))
	req.Header.Set("X-Legacy-Identity-Subject", req.Header.Get("X-Legacy-Identity-Subject"))
	req.Header.Set("X-Legacy-Auth-User", req.Header.Get("X-Legacy-Auth-User"))
	req.Header.Set("X-Legacy-Auth-Email", req.Header.Get("X-Legacy-Auth-Email"))
	req.Header.Set("X-Legacy-Auth-Preferred-Username", req.Header.Get("X-Legacy-Auth-Preferred-Username"))
	req.Header.Set("X-Legacy-Auth-Login", req.Header.Get("X-Legacy-Auth-Login"))
	req.Header.Set("X-Legacy-Auth-Principal", req.Header.Get("X-Legacy-Auth-Principal"))
	req.Header.Set("X-Legacy-Auth-Subject", req.Header.Get("X-Legacy-Auth-Subject"))
	req.Header.Set("X-Legacy-Remote-User", req.Header.Get("X-Legacy-Remote-User"))
	req.Header.Set("X-Legacy-Remote-Email", req.Header.Get("X-Legacy-Remote-Email"))
	req.Header.Set("X-Legacy-Remote-Preferred-Username", req.Header.Get("X-Legacy-Remote-Preferred-Username"))
	req.Header.Set("X-Legacy-Remote-Login", req.Header.Get("X-Legacy-Remote-Login"))
	req.Header.Set("X-Legacy-Remote-Principal", req.Header.Get("X-Legacy-Remote-Principal"))
	req.Header.Set("X-Legacy-Remote-Subject", req.Header.Get("X-Legacy-Remote-Subject"))
	req.Header.Set("X-Legacy-Session-User", req.Header.Get("X-Legacy-Session-User"))
	req.Header.Set("X-Legacy-Session-Email", req.Header.Get("X-Legacy-Session-Email"))
	req.Header.Set("X-Legacy-Session-Subject", req.Header.Get("X-Legacy-Session-Subject"))
}
