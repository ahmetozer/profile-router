package linkrouter

import (
	"github.com/ahmetozer/profile-router/config"
)

// FindProfile
// Find related profile for domain
func FindProfile(Config *config.Settings, link *string) string {
	for i := range Config.Site {
		if *link == Config.Site[i].Domain {
			return Config.Site[i].Profile
		}
	}
	return ""
}
