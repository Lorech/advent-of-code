package runners

import (
	"os/exec"
	"runtime"
)

// Opens the provided URL in the operating system's default browser.
func OpenURL(url string) error {
	var cmd, args = buildCommand(url)
	return exec.Command(cmd, args...).Start()
}

func buildCommand(url string) (string, []string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows": // Windows
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin": // macOS
		cmd = "open"
		args = []string{url}
	default: // Linux
		if isWSL() {
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			cmd = "xdg-open"
			args = []string{url}
		}
	}

	if len(args) > 1 {
		// args[0] is used for 'start' command argument
		// to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}

	return cmd, args
}
