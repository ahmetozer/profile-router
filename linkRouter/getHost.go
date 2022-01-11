package linkrouter

import (
	"net"
	"net/url"
	"strings"
)

// GetHost
// Get hostname from url
func GetHost(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	if strings.Contains(u.Host, ":") {
		host, _, _ := net.SplitHostPort(u.Host)
		if err != nil {
			return "", err
		}
		return host, nil
	}
	return u.Host, nil
}
