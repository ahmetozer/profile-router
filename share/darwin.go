//go:build darwin
// +build darwin

package share

const PATH_SEPARATOR string = "/"

var CHROME_LOCATIONS = []string{
	"/Applications/Chromium.app/Contents/MacOS/Chromium",
	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
}
