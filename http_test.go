package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectToHTTPSTrue(t *testing.T) {
	opts := NewOptions()
	opts.ForceHTTPS = true
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("test"))
	}

	h := redirectToHTTPS(opts, http.HandlerFunc(handler))
	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	h.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusPermanentRedirect, rw.Code, "status code should be %d, got: %d", http.StatusPermanentRedirect, rw.Code)
}

func TestRedirectToHTTPSFalse(t *testing.T) {
	opts := NewOptions()
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("test"))
	}

	h := redirectToHTTPS(opts, http.HandlerFunc(handler))
	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	h.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusOK, rw.Code, "status code should be %d, got: %d", http.StatusOK, rw.Code)
}

func TestRedirectNotWhenHTTPS(t *testing.T) {
	opts := NewOptions()
	opts.ForceHTTPS = true
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("test"))
	}

	h := redirectToHTTPS(opts, http.HandlerFunc(handler))
	s := httptest.NewTLSServer(h)
	defer s.Close()

	opts.HTTPSAddress = s.URL
	client := s.Client()
	res, err := client.Get(s.URL)
	if err != nil {
		t.Fatalf("request to test server failed with error: %v", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode, "status code should be %d, got: %d", http.StatusOK, res.StatusCode)
}
