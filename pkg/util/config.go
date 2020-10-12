package util

import (
	"fmt"
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
	configFilePath := getConfigPath()
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	confs := strings.Split(strings.Trim(strings.TrimSpace(string(content)), "\n"), "\n")
	if(confs == nil) {
		panic("Error splitting string")
	}
	configLen := len(confs)
	if configLen == 1 {
		return &HueConfig{
			Username: hueUser,
		}
	}
	if configLen == 2 {
		return &HueConfig{
			Username:  confs[0],
			IPAddress: confs[1],
		}
	}
	return nil
}

func SaveConfig(conf *HueConfig) {
	b := []byte(fmt.Sprintf("%s\n%s", conf.Username, conf.IPAddress))
	err := ioutil.WriteFile(getConfigPath(), b, 0600)
	if err != nil {
		panic("failed to write config to ~/.hue")
	}
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, ".hue")
}
