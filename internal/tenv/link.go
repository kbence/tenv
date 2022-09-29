package tenv

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func ExistsOnPath(binaryName string) (bool, error) {
	binPath := os.Getenv("PATH")
	binPathDirs := strings.Split(binPath, string(os.PathListSeparator))

	for _, dir := range binPathDirs {
		stat, err := os.Stat(path.Join(dir, binaryName))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return false, err
		}

		if stat.Mode().Perm()&0111 != 0 {
			return true, nil
		}
	}

	return false, nil
}

func CreateSymbolicLink(binaryName string) error {
	tenvBinaryPath, err := os.Executable()
	if err != nil {
		return err
	}

	tenvBinaryDir, tenvBinaryName := path.Split(tenvBinaryPath)

	return os.Symlink(tenvBinaryName, path.Join(tenvBinaryDir, binaryName))
}

func CreateLinks(force bool) error {
	if !force {
		for _, binaryName := range TeleportBinaryNames {
			if exists, err := ExistsOnPath(binaryName); err != nil {
				return fmt.Errorf("error checking for '%s' on PATH: %s", binaryName, err)
			} else if exists {
				return fmt.Errorf("binary '%s' already exists on PATH", binaryName)
			}
		}
	}

	for _, binaryName := range TeleportBinaryNames {
		err := CreateSymbolicLink(binaryName)
		if err != nil {
			return err
		}
	}

	return nil
}
