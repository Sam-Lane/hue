package util

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

//HueConfig stores configurations for the use with a Hue Bridge
type HueConfig struct {
	IPAddress string
	Username  string
}

func ReadConfig() *HueConfig {
	var hueUser string

	//Check if user has set hue username in the env vars.
	hueUser, exists := os.LookupEnv("HUE_USERNAME")
	if exists {
		return &HueConfig{
			Username: hueUser,
		}
	}
	//If not proceed to read .hue file from their home directory

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(usr.HomeDir, ".hue")

	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	hueUser = strings.TrimSpace(string(content))
	return &HueConfig{
		Username: hueUser,
	}
}
