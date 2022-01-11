package linkrouter

import (
	"os"

	"github.com/ahmetozer/profile-router/config"
)

// ProfileExecuter
// Execute chrome with selected profile

func ProfileExecuter(Config *config.Settings, profile string) {
	profileArg := "--profile-directory=" + profile
	var args []string
	args = append(args, profileArg)
	if len(os.Args) != 0 {
		args = append(args, os.Args[1:]...)
	}
	Executer(Config, &args)
}
