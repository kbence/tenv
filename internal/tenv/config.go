package tenv

import (
	"os"
	"path"

	"github.com/go-yaml/yaml"
)

const FileMode = 0755
const DirectoryMode = 0755

var TeleportEnvHomeDirectory = os.ExpandEnv("$HOME/.tenv")
var TeleportEnvConfigFile = path.Join(TeleportEnvHomeDirectory, "config.yaml")
var TeleportEnvVersionDirectory = path.Join(TeleportEnvHomeDirectory, "versions")
var TeleportHomeDirectory = os.ExpandEnv("$HOME/.tsh")
var TeleportCurrentProfileFile = path.Join(TeleportHomeDirectory, "current-profile")

var TeleportBinaryNames = []string{"teleport", "tsh", "tctl", "tbot"}
var TeleportOptionalBinaryNames = []string{"tbot"}

type Config struct {
	Rules []struct {
		Match []struct {
			CurrentDir     *string `yaml:"currentDir,omitempty"`
			CurrentProfile *string `yaml:"currentProfile,omitempty"`
		} `yaml:"match,omitempty"`
		Version string `yaml:"version,omitempty"`
	} `yaml:"rules,omitempty"`
}

func LoadConfig() (*Config, error) {
	var config Config

	content, err := os.ReadFile(TeleportEnvConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &config, nil
		}
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
