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
	_ = forwardedProto
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
	trustedControlValue := "signed-internal-request"
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
	req.Header.Set("X-Legacy-Identity-User", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Identity-Email", legacyIdentityValue(authenticatedEmail(session)))
	req.Header.Set("X-Legacy-Identity-Preferred-Username", legacyIdentityValue(session.PreferredUsername))
	req.Header.Set("X-Legacy-Identity-Login", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Identity-Principal", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Identity-Subject", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Auth-User", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Auth-Email", legacyIdentityValue(authenticatedEmail(session)))
	req.Header.Set("X-Legacy-Auth-Preferred-Username", legacyIdentityValue(session.PreferredUsername))
	req.Header.Set("X-Legacy-Auth-Login", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Auth-Principal", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Auth-Subject", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Remote-User", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Remote-Email", legacyIdentityValue(authenticatedEmail(session)))
	req.Header.Set("X-Legacy-Remote-Preferred-Username", legacyIdentityValue(session.PreferredUsername))
	req.Header.Set("X-Legacy-Remote-Login", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Remote-Principal", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Remote-Subject", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Session-User", legacyIdentityValue(authenticatedUser(session)))
	req.Header.Set("X-Legacy-Session-Email", legacyIdentityValue(authenticatedEmail(session)))
	req.Header.Set("X-Legacy-Session-Subject", legacyIdentityValue(authenticatedUser(session)))
}
