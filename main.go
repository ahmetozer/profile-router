package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/ahmetozer/profile-router/config"
	"github.com/ahmetozer/profile-router/gui"
	linkrouter "github.com/ahmetozer/profile-router/linkRouter"
)

const (
	APP_NAME = "Link Router"
)

var (
	Config      config.Settings
	ConfigError error
)

func main() {

	Config, ConfigError = config.LoadConfig()
	linkrouter.LookRouteTable(&Config)
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title(APP_NAME),
			app.Size(unit.Dp(200), unit.Dp(300)),
		)
		if err := gui.Window(w, &Config, &ConfigError); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()

}
