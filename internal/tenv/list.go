package tenv

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type version string

func (v version) Major() int    { return v.VersionPart(0) }
func (v version) Minor() int    { return v.VersionPart(1) }
func (v version) Revision() int { return v.VersionPart(2) }

func (v version) VersionPart(n int) int {
	parts := strings.Split(v.Version(), ".")

	if len(parts) <= n {
		return 0
	}

	ver, _ := strconv.Atoi(parts[n])
	return ver
}

func (v version) Version() string {
	parts := strings.SplitN(string(v), "-", 2)
	return parts[0]
}

func (v version) Tag() string {
	parts := strings.SplitN(string(v), "-", 2)

	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}

type versionList []version

func (l versionList) Len() int {
	return len(l)
}

func (l versionList) Less(a, b int) bool {
	for i := 0; i < 3; i++ {
		if l[a].VersionPart(i) < l[b].VersionPart(i) {
			return true
		}
		if l[a].VersionPart(i) > l[b].VersionPart(i) {
			return false
		}
	}

	return l[a].Tag() < l[b].Tag()
}

func (l versionList) Swap(a, b int) {
	l[a], l[b] = l[b], l[a]
}

func (l versionList) StringSlice() []string {
	newSlice := []string{}

	for _, v := range l {
		newSlice = append(newSlice, string(v))
	}

	return newSlice
}

func ContainsTeleportBinaries(dir string) bool {
	binDir := path.Join(dir, "bin")
	stat, err := os.Stat(binDir)
	if err != nil {
		return false
	}

	if !stat.IsDir() {
		return false
	}

	entries, err := os.ReadDir(binDir)
	if err != nil {
		return false
	}

	missingBinaries := map[string]struct{}{}
	for _, binary := range TeleportBinaryNames {
		if !StringSliceContains(TeleportOptionalBinaryNames, binary) {
			missingBinaries[binary] = struct{}{}
		}
	}

	for _, entry := range entries {
		delete(missingBinaries, entry.Name())
	}

	return len(missingBinaries) == 0
}

func GetInstalledVersions() (versionList, error) {
	entries, err := os.ReadDir(TeleportEnvVersionDirectory)
	if err != nil {
		return nil, err
	}

	versions := versionList{}

	for _, entry := range entries {
		versionDir := path.Join(TeleportEnvVersionDirectory, entry.Name())
		if ContainsTeleportBinaries(versionDir) {
			versions = append(versions, version(entry.Name()))
		}
	}

	sort.Sort(versions)

	return versions, nil
}

func ListInstalledVersions() error {
	versions, err := GetInstalledVersions()
	if err != nil {
		return err
	}

	for _, ver := range versions {
		fmt.Println(ver)
	}

	return nil
}
