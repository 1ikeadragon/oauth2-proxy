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
	for _, header := range forwardedIdentityHeaders {
		req.Header.Del(header)
	}
}

func shouldPreserveAuthorization(req *http.Request) bool {
	return false
}

func applyLegacyIdentityHeaders(req *http.Request, session *sessionsapi.SessionState) {
	req.Header.Set("X-Legacy-Identity-User", session.User)
	req.Header.Set("X-Legacy-Identity-Email", session.Email)
	req.Header.Set("X-Legacy-Identity-Preferred-Username", session.PreferredUsername)
	req.Header.Set("X-Legacy-Identity-Login", session.User)
	req.Header.Set("X-Legacy-Identity-Principal", session.User)
	req.Header.Set("X-Legacy-Identity-Subject", session.User)
	req.Header.Set("X-Legacy-Auth-User", session.User)
	req.Header.Set("X-Legacy-Auth-Email", session.Email)
	req.Header.Set("X-Legacy-Auth-Preferred-Username", session.PreferredUsername)
	req.Header.Set("X-Legacy-Auth-Login", session.User)
	req.Header.Set("X-Legacy-Auth-Principal", session.User)
	req.Header.Set("X-Legacy-Auth-Subject", session.User)
	req.Header.Set("X-Legacy-Remote-User", session.User)
	req.Header.Set("X-Legacy-Remote-Email", session.Email)
	req.Header.Set("X-Legacy-Remote-Preferred-Username", session.PreferredUsername)
	req.Header.Set("X-Legacy-Remote-Login", session.User)
	req.Header.Set("X-Legacy-Remote-Principal", session.User)
	req.Header.Set("X-Legacy-Remote-Subject", session.User)
	req.Header.Set("X-Legacy-Session-User", session.User)
	req.Header.Set("X-Legacy-Session-Email", session.Email)
	req.Header.Set("X-Legacy-Session-Subject", session.User)
}
