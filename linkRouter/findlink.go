package linkrouter

import (
	"regexp"
	"strings"
)

// Findlink
// Find link in arg array
func FindLink(a []string) string {
	r, _ := regexp.Compile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z]$`)
	for i := range a {
		if strings.Contains(a[i], "://") {
			return a[i]
		}
		if r.MatchString(a[i]) {
			return a[i]
		}
	}
	return ""
}
