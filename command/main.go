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

func main() {
	if len(os.Args) == 1 {
		displayHelp()
		os.Exit(0)
	}

	shouldCreatePackage, stagingOverride, err := getArguments()
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

	if shouldCreatePackage {
		fmt.Println("Creating package...")
		path, err := pacchetto.CreatePackage(destinationParentPath, stagingOverride)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Created package at '" + path + "'")
	}
}

func getArguments() (shouldCreatePackage bool, stagingOverride string, err error) {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		return false, "", errors.New("You must specify an argument")
	}

	// Super hack to detect if the string is an option, or an argument.
	var hasOption bool
	for _, temp := range arguments {
		if strings.HasPrefix(temp, "-") {
			hasOption = true
			break
		}
	}
	if !hasOption {
		return false, "", errors.New("You must specify an option")
	}

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
			case "p":
				shouldCreatePackage = true
			case "s":
				stagingOverride = value
			default:
				return false, "", errors.New("Unknown option: '" + argument + "'")
			}
		}
	}

	return shouldCreatePackage, stagingOverride, nil
}

func displayHelp() {
	// TODO: This is terrible. Do something better.
	lines := []string{
		"usage: pacchetto",
		"",
		"Options:",
		"    -h    Display this help page.",
		"    -p    Create a package.",
		"    -s    The staging path override flag. This allows you to override the.",
		"          temporaary directory where files are staged.",
		"          This is useful if your computer has limited storage space",
		"          on the main disk.",
		"",
		"Examples:",
		"    'pacchetto -p'",
		"    Create a single archive containing all the files needed to run an Assetto",
		"    Corsa dedicated server.",
		"",
		"    'pacchetto -p -s \"E:/alternative/temp/dir\"'",
		"    Create a single archive containing all the files needed to run an Assetto",
		"    Corsa dedicated server using an alternative directory for staging files.",
		"",
	}

	lineEnding := "\n"
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	}

	fmt.Println(strings.Join(lines, lineEnding))
}
