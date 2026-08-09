package diagnostic

import "os/exec"

func Run(target string) ([]byte, error) {
	return exec.Command("sh", "-c", "ping -c 1 "+target).CombinedOutput()
}