package tenv

import (
	"fmt"
	"os"
	"path"
)

func GetValidProfileNames() ([]string, error) {
	entries, err := os.ReadDir(path.Join(TeleportHomeDirectory, "keys"))
	if err != nil {
		return nil, err
	}

	var profileNames []string

	for _, entry := range entries {
		if entry.IsDir() {
			profileNames = append(profileNames, entry.Name())
		}
	}

	return profileNames, nil
}

func SelectProfile(profileName string) error {
	return os.WriteFile(TeleportCurrentProfileFile, []byte(fmt.Sprintf("%s\n", profileName)), 0644)
}
