package tenv

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"
)

var TeleportBinaryURLTemplate = template.Must(template.New("binary_url").Parse(
	"https://get.gravitational.com/teleport-v{{.Version}}-{{.OS}}-{{.Arch}}-bin.tar.gz",
))

type versionDescriptor struct {
	OS      string
	Arch    string
	Version string
}

func downloadTeleportBinaries(ctx context.Context, targetDir string, version string) error {
	err := os.MkdirAll(targetDir, DirectoryMode)
	if err != nil {
		return err
	}

	urlWriter := &strings.Builder{}
	err = TeleportBinaryURLTemplate.Execute(urlWriter, versionDescriptor{
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
		Version: version,
	})
	if err != nil {
		return err
	}

	req, err := http.Get(urlWriter.String())
	if err != nil {
		return err
	}

	gzReader, err := gzip.NewReader(req.Body)
	if err != nil {
		return err
	}

	reader := tar.NewReader(gzReader)

	for {
		header, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		dir, file := path.Split(header.Name)
		if dir != "teleport/" {
			continue
		}

		switch file {
		case "tbot", "teleport", "tsh", "tctl":
			if !header.FileInfo().Mode().IsRegular() {
				return fmt.Errorf("file expected to be regular: %s", header.Name)
			}

			f, err := os.OpenFile(path.Join(targetDir, file), os.O_CREATE|os.O_WRONLY, header.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			io.Copy(f, reader)
			log.Printf("Copied %s to %s\n", header.Name, path.Join(targetDir, file))
		}
	}

	return nil
}

func BinDirectory(version string) string {
	return path.Join(TeleportEnvVersionDirectory, version, "bin")
}

func InstallVersion(ctx context.Context, version string) error {
	err := downloadTeleportBinaries(ctx, BinDirectory(version), version)
	if err != nil {
		return err
	}

	return nil
}
