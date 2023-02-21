package tenv

import (
	"fmt"
	"os"
	"path"
	"syscall"
)

var SelectedVersionFile = path.Join(TeleportEnvHomeDirectory, "selected-version")

func UseTeleport(version string) error {
	err := os.MkdirAll(TeleportEnvHomeDirectory, DirectoryMode)
	if err != nil {
		return err
	}

	err = os.WriteFile(SelectedVersionFile, []byte(fmt.Sprintf("%s\n", version)), FileMode)
	if err != nil {
		return err
	}

	return nil
}

func Execute(binaryName string, args ...string) error {
	version, err := GetSelectedVersion()
	if err != nil {
		return err
	}

	binary := path.Join(BinDirectory(version), binaryName)

	return syscall.Exec(binary, args, os.Environ())
}
