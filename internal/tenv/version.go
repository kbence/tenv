package tenv

import (
	"os"
	"strings"
)

func getVersionFromRules() (*string, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	if len(config.Rules) == 0 {
		return nil, nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	currentProfileContent, err := os.ReadFile(TeleportCurrentProfileFile)
	if err != nil {
		return nil, err
	}
	currentProfile := strings.TrimSpace(string(currentProfileContent))

	for _, rule := range config.Rules {
		for _, match := range rule.Match {
			matchedAll := true

			if match.CurrentDir != nil {
				matchedAll = matchedAll && strings.HasPrefix(currentDir, *match.CurrentDir)
			}

			if match.CurrentProfile != nil {
				matchedAll = matchedAll && strings.TrimSpace(*match.CurrentProfile) == currentProfile
			}

			if matchedAll {
				return &rule.Version, nil
			}
		}
	}

	return nil, err
}

func GetSelectedVersion() (string, error) {
	versionFromEnv := os.Getenv("TELEPORT_VERSION")
	if len(versionFromEnv) > 0 {
		return versionFromEnv, nil
	}

	versionFromRules, err := getVersionFromRules()
	if err != nil {
		return "", err
	}

	if versionFromRules != nil {
		return *versionFromRules, nil
	}

	content, err := os.ReadFile(SelectedVersionFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
