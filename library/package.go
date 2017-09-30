package pacchetto

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/mholt/archiver"
	"github.com/stephen-fox/cabinet"
	"github.com/stephen-fox/logi"
)

// CreatePackage creates a single archive in the specified parent directory
// that contains all of the files required to run an Assetto Corsa dedicated
// server. Optionally, the caller may override the parent path of the
// staging directory.
func CreatePackage(parentDirPath string, stagingPathOverride string) (archivePath string, err error) {
	acPath, err := GetAssettoCorsaPath()
	if err != nil {
		return "", err
	}

	// If the staging path override is not specified, then the OS' temp dir is
	// used instead.
	tempDirPath, err := ioutil.TempDir(stagingPathOverride, tempPrefix+".")
	if err != nil {
		return "", errors.New("Failed to create temporary directory")
	}
	defer os.RemoveAll(tempDirPath)

	logi.Info.Println("Staging server files...")
	serverStagingPath := tempDirPath + "/" + serverSubPath
	err = cabinet.CopyDirectory(acPath+"/"+serverSubPath, serverStagingPath, true)
	if err != nil {
		return "", err
	}

	for _, subPath := range contentSubPaths {
		path := acPath + "/content/" + subPath
		if !cabinet.Exists(path) {
			return "", errors.New("Assetto Corsa content directory '" + path +
				"' does not exist")
		}

		logi.Info.Println("Staging content", path, "...")

		err := cabinet.CopyFilesWithSuffix(path, serverStagingPath+"/content/"+subPath, ".ini", true)
		if err != nil {
			return "", errors.New("Failed to stage '" + path + "'")
		}

		err = cabinet.CopyFilesWithSuffix(path, serverStagingPath+"/content/"+subPath, ".acd", true)
		if err != nil {
			return "", errors.New("Failed to stage '" + path + "'")
		}
	}

	logi.Info.Println("Creating package...")
	archiveDirs := []string{
		serverStagingPath,
	}
	fullDestinationPath := parentDirPath + "/" + outputPrefix + ".zip"
	err = archiver.Zip.Make(fullDestinationPath, archiveDirs)
	if err != nil {
		return "", err
	}

	logi.Info.Println("Successfully created server package")
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
			tempFilePath := l + ":/.pacchetto"
			temp, err := os.Create(tempFilePath)
			if err != nil {
				continue
			}
			temp.Close()
			os.Remove(temp.Name())
			path = l + subPath
			if cabinet.Exists(path) {
				return path, nil
			}
		}
	}

	return "", errors.New("Failed to locate Assetto Corsa directory")
}
