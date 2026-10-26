package plantedvuln

import "os/exec"

// Planted test vulnerability: OS command injection via unsanitized input.
func Ping(host string) ([]byte, error) {
	return exec.Command("sh", "-c", "ping -c 1 "+host).CombinedOutput()
}
