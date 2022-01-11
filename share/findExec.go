package share

import (
	"fmt"
	"os/exec"
)

// FindExecPath
// Try to find google chrome
func FindExecPath() (string, error) {
	for _, path := range CHROME_LOCATIONS {
		found, err := exec.LookPath(path)
		if err == nil {
			return found, nil
		}
	}
	return "", fmt.Errorf("chrome path is not found")
}
