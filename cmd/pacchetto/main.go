package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/stephen-fox/pacchetto"
)

const (
	applicationName = "pacchetto"

	shouldCreatePackageArg = "p"
	stagingPathOverrideArg = "s"
	shouldPrintHelpArg     = "h"
	shouldPrintVersionArg  = "v"
	shouldPrintExamplesArg = "x"

	examples = `	Create an Assetto Corsa dedicated server pacakge that can be distributed
	to other machines:
		` + applicationName + ` -` + shouldCreatePackageArg + `

	Create an Assetto Corsa dedicated server pacakge using an alternative
	staging directory. This is useful if your main storage is full:
		` + applicationName + `-` + shouldCreatePackageArg + ` -` + stagingPathOverrideArg + ` E:/alternative/temp/dir`
)

var (
	version string

	stagingPathOverride = flag.String(stagingPathOverrideArg, "", "Optionally override the path where the server files are staged")

	shouldCreatePackage = flag.Bool(shouldCreatePackageArg, false, "Create an Assetto Corsa dedicated server package")
	shouldPrintHelp     = flag.Bool(shouldPrintHelpArg, false, "Prints this help page")
	shouldPrintExamples = flag.Bool(shouldPrintExamplesArg, false, "Print application usage examples")
	shouldPrintVersion  = flag.Bool(shouldPrintVersionArg, false, "Print the application version")
)

func main() {
	flag.Parse()

	if *shouldPrintHelp || len(os.Args) == 1 {
		fmt.Println(applicationName, version)
		fmt.Println()
		fmt.Println("[ABOUT]")
		fmt.Println("Tool for creating Assetto Corsa dedicated server packages, which can be directly")
		fmt.Println("distributed to other machines.")
		fmt.Println()
		fmt.Println("[USAGE]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *shouldPrintExamples {
		fmt.Println(examples)
		os.Exit(0)
	}

	if *shouldPrintVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *shouldCreatePackage {
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

		log.Println("Creating package...")
		path, err := pacchetto.CreatePackage(destinationParentPath, *stagingPathOverride)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("Created package at '" + path + "'")
	}
}
