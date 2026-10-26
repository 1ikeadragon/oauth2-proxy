package plantedvuln

import (
	"crypto/tls"
	"net/http"
)

// Planted test vulnerability: TLS certificate verification disabled.
func NewInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
