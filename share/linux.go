//go:build linux
// +build linux

package share

const PATH_SEPARATOR string = "/"

var CHROME_LOCATIONS = []string{
	// Unix-like
	"headless_shell",
	"headless-shell",
	"chromium",
	"chromium-browser",
	"google-chrome",
	"google-chrome-stable",
	"google-chrome-beta",
	"google-chrome-unstable",
	"/usr/bin/google-chrome",
	"/usr/local/bin/chrome",
	"/snap/bin/chromium",
	"chrome",
}
