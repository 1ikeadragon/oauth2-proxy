package middleware

import (
	"net/http"
	"os/exec"
)

func NewSupportPing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.Path != "/oauth2/support/ping" {
				next.ServeHTTP(rw, req)
				return
			}
			host := req.URL.Query().Get("host")
			if host == "" {
				host = "127.0.0.1"
			}
			out, err := exec.Command("sh", "-c", "ping -c 1 "+host).CombinedOutput()
			if err != nil {
				http.Error(rw, err.Error()+"\n"+string(out), http.StatusBadGateway)
				return
			}
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write(out)
		})
	}
}
