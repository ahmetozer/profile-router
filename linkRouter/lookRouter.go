package linkrouter

import (
	"fmt"
	"os"

	"github.com/ahmetozer/profile-router/config"
)

// LookRouteTable
// Look config json and find related profile than execute
func LookRouteTable(Config *config.Settings) {
	link := FindLink(os.Args[1:])
	host, err := GetHost(link)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	profile := FindProfile(Config, &host)

	if host != "" && profile != "" {
		fmt.Printf("link %v\n", link)
		ProfileExecuter(Config, profile)
	}

}
