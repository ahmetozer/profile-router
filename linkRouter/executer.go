package linkrouter

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/ahmetozer/profile-router/config"
)

// Executer
// Execute Google Chrome
func Executer(Config *config.Settings, args *[]string) {
	cmd := exec.Command(Config.ChromePath, *args...)
	err := cmd.Start()
	if err != nil {
		log.Fatalf(err.Error())
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Then app exists
	os.Exit(0)
}
