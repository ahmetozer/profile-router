package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ahmetozer/profile-router/share"
)

// LoadConfig
// Load configuration from json
func LoadConfig() (Settings, error) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("dirname: %v", err)
	}
	data := Settings{}
	var config_dir = dirname + share.PATH_SEPARATOR + ".profilerouter.json"
	file, err := ioutil.ReadFile(config_dir)

	if err != nil {
		return data, fmt.Errorf("err while opening config file,%v", err)
	}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		return data, fmt.Errorf("err while parsing config file,%v", err)
	}

	if data.ChromePath == "" {
		data.ChromePath, err = share.FindExecPath()
	}
	return data, err
}
