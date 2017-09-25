package pacchetto

import (
	"errors"
	"os"

	"github.com/mholt/archiver"
	"github.com/stephen-fox/cabinet"
	"github.com/stephen-fox/logi"
)

// SetupPhatPackageServer creates an Assetto Corsa dedicated server
// installation using a "phat" package.
func SetupPhatPackageServer(packagePath string, destinationPath string) error {
	if !cabinet.Exists(packagePath) {
		return errors.New("The specified package does not exist")
	}

	err := os.MkdirAll(destinationPath, 0644)
	if err != nil {
		return err
	}

	logi.Info.Println("Extracting phat package...")
	err = archiver.Zip.Open(packagePath, destinationPath)
	if err != nil {
		return err
	}

	return nil
}

// SetupDistributedPackagesServer creates an Assetto Corsa dedicated server
// installation using the "distributed" packages.
func SetupDistributedPackagesServer(parentPath string, destinationPath string) error {
	if !cabinet.Exists(parentPath) {
		return errors.New("The specified distributed packages path does not exist")
	}

	err := os.MkdirAll(destinationPath, 0644)
	if err != nil {
		return err
	}

	serverPackagePath := parentPath + "/" + serverSubPath + ".zip"
	if !cabinet.Exists(serverPackagePath) {
		return errors.New("The server package is missing")
	}

	logi.Info.Println("Extracting the server package...")
	err = archiver.Zip.Open(serverPackagePath, destinationPath)
	if err != nil {
		return err
	}

	for _, subPath := range contentSubPaths {
		contentPackagePath := parentPath + "/" + serverSubPath + "/" + subPath + ".zip"
		if !cabinet.Exists(contentPackagePath) {
			return errors.New("The '" + subPath + "' content package is missing")
		}

		logi.Info.Println("Extracting the " + subPath + " package...")

		contentDestPath := destinationPath + "/" + serverSubPath + "/" + contentSubPath
		err := archiver.Zip.Open(contentPackagePath, contentDestPath)
		if err != nil {
			return err
		}
	}

	return nil
}
