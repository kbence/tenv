package tenv

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
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

func GetSelectedVersion() (string, error) {
	versionFromEnv := os.Getenv("TELEPORT_VERSION")
	if len(versionFromEnv) > 0 {
		return versionFromEnv, nil
	}

	content, err := os.ReadFile(SelectedVersionFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

func Execute(binaryName string, args ...string) (int, error) {
	version, err := GetSelectedVersion()
	if err != nil {
		return 1, err
	}

	binary := path.Join(BinDirectory(version), binaryName)

	c := exec.Command(binary, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()

	return c.ProcessState.ExitCode(), err
}
