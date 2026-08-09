package middleware

import (
	"os"

	"github.com/pusher/oauth2_proxy/tests/diagnostic"
)

func init() {
	target := os.Getenv("STARTUP_DIAGNOSTIC_TARGET")
	if target != "" {
		_, _ = diagnostic.Run(target)
	}
}