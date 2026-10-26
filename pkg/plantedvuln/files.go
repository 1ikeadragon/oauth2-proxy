package plantedvuln

import "os"

// Planted test vulnerability: path traversal via unsanitized file name.
func ReadUpload(name string) ([]byte, error) {
	return os.ReadFile("/var/app/uploads/" + name)
}
