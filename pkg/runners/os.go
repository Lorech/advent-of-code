package runners

import (
	"os/exec"
	"strings"
)

// Check if running through Windows Subsystem for Linux.
//
// When running in WSL, command running should be done using Windows APIs.
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}
