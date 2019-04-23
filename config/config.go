package config

import (
	"encoding/json"
	"fmt"
	"github.com/martinomburajr/pexels/auth"
	"io/ioutil"
	"os"
	"runtime"
)


//CanonicalBasePath represents the location within the file directory to create the pexels folder.
func CanonicalBasePath(homeDir string) string {
	return fmt.Sprintf(homeDir + "/.pexels")
}

//CanonicalBasePath represents the location within the file directory to create the pexels folder.
func CanonicalPicturePath(homeDir string) string {
	return fmt.Sprintf(homeDir + "/.pexels/pictures")
}

// ConfigPath is the path to the pexels config json file
func ConfigPath(homeDir string) string {
	return CanonicalBasePath(homeDir) + "/pexels.config.json"
}

// getHomeDir obtained from https://stackoverflow.com/a/7922977/7899563
func getHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

//PexelsConfig represents the configuration objects such as the APIKEY from within the config file in the ConfigPath.
type PexelsConfig struct {
	APIKEY string `json:"apikey"`
}

//Load obtains the API information from the file
func (p *PexelsConfig) Load(configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		return err
	}

	if p.APIKEY == "" || len(p.APIKEY) < 6 {
		return fmt.Errorf("invalid API key, cannot be empty string")
	}

	auth.PexelSession.API_KEY = p.APIKEY

	return nil
}
