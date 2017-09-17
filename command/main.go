package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/stephen-fox/pacchetto/library"
)

func main() {
	destinationParentPath := ""
	switch operatingSystem := runtime.GOOS; operatingSystem {
	case "darwin":
		destinationParentPath = os.Getenv("HOME")
	case "linux":
		destinationParentPath = os.Getenv("HOME")
	case "windows":
		destinationParentPath = filepath.ToSlash(os.Getenv("USERPROFILE"))
	}
	destinationParentPath = destinationParentPath + "/Desktop"

	destination, err := pacchetto.PackageAssettoCorsaServer(destinationParentPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Assetto Corsa archive created at: '" + destination + "'")
}
