package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/stephen-fox/pacchetto/library"
)

type mode int

const (
	phat mode = iota
	distributed
)

func main() {
	if len(os.Args) == 1 {
		displayHelp()
		os.Exit(0)
	}

	mode, stagingOverride, err := getArguments()
	if err != nil {
		fmt.Println(err.Error(), "- Re-run with the '-h' argument for more information")
		os.Exit(1)
	}

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

	if mode == distributed {
		fmt.Println("Creating packages...")
		path, err := pacchetto.CreateDistributedPackages(destinationParentPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Created packages at " + path)
	} else if mode == phat {
		fmt.Println("Creating package...")
		path, err := pacchetto.CreatePhatPackage(destinationParentPath, stagingOverride)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Created package at " + path)
	}
}

func getArguments() (mode mode, stagingOverride string, err error) {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return 0, "", errors.New("You must specify an argument")
	}

	// Super hack.
	var hasOption bool
	for _, temp := range arguments {
		if strings.HasPrefix(temp, "-") {
			hasOption = true
			break
		}
	}
	if !hasOption {
		return 0, "", errors.New("You must specify an option")
	}

	var hasMode bool
	for i, argument := range arguments {
		value := ""
		if i < len(arguments)-1 {
			// There is another argument (element) available.
			// Add two because we have skipped element 0, and the value is the
			// next element in the slice.
			value = os.Args[i+2]
		}
		if strings.HasPrefix(argument, "-") || strings.HasPrefix(argument, "--") {
			parsedArgument := strings.TrimPrefix(argument, "-")
			parsedArgument = strings.TrimPrefix(argument, "-")
			if len(parsedArgument) == 0 {
				continue
			}
			switch parsedArgument {
			case "h":
				displayHelp()
				os.Exit(0)
			case "m":
				hasMode = true
				if value == "distributed" || value == "d" {
					mode = distributed
				} else if value == "phat" || value == "p" {
					mode = phat
				} else {
					return 0, "", errors.New("Mode " + value + " is invalid")
				}
			case "s":
				stagingOverride = value
			default:
				return 0, "", errors.New("Unknown option: '" + argument + "'")
			}
		}
	}

	if !hasMode {
		return 0, "", errors.New("Missing 'mode' argument")
	}

	return mode, stagingOverride, nil
}

func displayHelp() {
	// TODO: This is terrible. Do something better.
	lines := []string{
		"usage: pacchetto",
		"",
		"Options:",
		"    -h    Display this help page.",
		"    -m    The mode flag. This can be set to 'distributed' or 'phat'.",
		"    -s    The staging path override flag. This allows you to override the.",
		"          temporaary directory where the 'phat' archive files are staged.",
		"          This may be needed if your computer has limited storage space",
		"          on the main disk.",
		"",
		"Examples:",
		"    'pacchetto -m phat'",
		"    Create a single archive containing all the files needed to run an Assetto",
		"    Corsa dedicated server.",
		"",
		"    'pacchetto -m phat -s E:/alternative/temp/dir'",
		"    Create a single archive containing all the files needed to run an Assetto",
		"    Corsa dedicated server using an alternative directory for staging files.",
		"",
		"    'pacchetto -m distributed'",
		"    Create several archives, each containing certain files required to run",
		"    Assetto Corsa dedicated server. This may be preferable over a single",
		"    massive archive file.",
	}

	lineEnding := "\n"
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	}

	fmt.Println(strings.Join(lines, lineEnding))
}
