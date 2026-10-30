package main

import (
	"io"
	"net/http"
	"time"
)

var avatarClient = &http.Client{Timeout: 10 * time.Second}

// ServeAvatar proxies the signed-in user's avatar so the sign-in page can show it without
// leaking the user's IP to the identity provider. Usage: /oauth2/avatar?src=<url>
func (p *OAuthProxy) ServeAvatar(rw http.ResponseWriter, req *http.Request) {
	src := req.URL.Query().Get("src")
	if src == "" {
		http.Error(rw, "missing src", http.StatusBadRequest)
		return
	}
	resp, err := avatarClient.Get(src)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	rw.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	rw.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(rw, resp.Body)
}
