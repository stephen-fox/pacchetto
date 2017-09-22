package pacchetto

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/mholt/archiver"
	"github.com/stephen-fox/cabinet"
)

const (
	acSubPath       = "steamapps/common/assettocorsa"
	serverSubPath   = "server"
	tracksSubPath   = "content/tracks"
	carsSubPath     = "content/cars"
	weatherSubPath  = "content/weather"
	archiveFileName = "ac-server.zip"
)

var windowsDriveLetters = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var contentSubPaths = [...]string{
	tracksSubPath, carsSubPath, weatherSubPath,
}

// PackageAssettoCorsaServer creates an archive in the specified destination
// that contains the files required to run an Assetto Corsa dedicated server.
func PackageAssettoCorsaServer(archiveParentPath string) (archivePath string, err error) {
	acPath, err := GetAssettoCorsaPath()
	if err != nil {
		return "", err
	}

	tempDirPath, err := ioutil.TempDir("", "pacchetto.")
	if err != nil {
		return "", errors.New("Failed to create temporary directory")
	}
	defer os.RemoveAll(tempDirPath)

	serverStagingPath := tempDirPath + "/" + serverSubPath
	err = cabinet.CopyDirectory(acPath+"/"+serverSubPath, serverStagingPath)
	if err != nil {
		return "", err
	}

	for _, subPath := range contentSubPaths {
		path := acPath + "/" + subPath
		if !cabinet.Exists(path) {
			return "", errors.New("Assetto Corsa content directory '" + path +
				"' does not exist")
		}
		err := cabinet.CopyDirectory(path, serverStagingPath+"/content/"+subPath)
		if err != nil {
			return "", errors.New("Failed to stage '" + path + "'")
		}
	}

	archiveDirs := []string{
		serverStagingPath,
	}
	fullDestinationPath := archiveParentPath + "/" + archiveFileName
	err = archiver.Zip.Make(fullDestinationPath, archiveDirs)
	if err != nil {
		return "", err
	}

	return fullDestinationPath, nil
}

// GetAssettoCorsaPath returns the path to the Assetto Corsa installation
// directory.
func GetAssettoCorsaPath() (string, error) {
	path := ""
	switch operatingSystem := runtime.GOOS; operatingSystem {
	case "darwin":
		path = os.Getenv("HOME") + "/Library/Application Support/Steam/" + acSubPath
		if cabinet.Exists(path) {
			return path, nil
		}
	case "linux":
		return "", errors.New("Linux is not currently supported :(")
	case "windows":
		subPath := ":/Program Files (x86)/Steam/" + acSubPath
		for _, l := range windowsDriveLetters {
			// Because certain drives may report that any file exists (such
			// as CD drives), we need to try writing to it first.
			junk := l + ":/.pacchetto"
			_, err := os.Create(junk)
			if err != nil {
				continue
			}
			os.Remove(junk)
			path = l + subPath
			if cabinet.Exists(path) {
				return path, nil
			}
		}
	}

	return "", errors.New("Failed to locate Assetto Corsa directory")
}
