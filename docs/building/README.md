# Build documentation
This document describes how to build the application.

## Requirements
In order to compile the application, you will need the following:
* [Java](http://www.oracle.com/technetwork/java/javase/downloads/)
* [Go](https://golang.org/dl/)

## Setting up the build environment
The build environment for the application requires additional environment
variables and specific Go libraries. The following sections will guide you
through this.

### Setting up environment variables
First, you must create the `GOPATH` environment variable. This should be set
to `/path/to/your/home/go`. You can then add `GOPATH/bin` to your `PATH`.
Lastly, you will want to add `/path/to/go/bin` to your `PATH`.

### Setting up required Go libraries
In addition to the above tools, you will also need to acquire several Go
libraries. This is automated via a Gradle task. To run this task, execute the
following:
```
./gradlew setupGoLibraries
```

## Building for Linux, macOS, and Windows
Execute the following on the command line:
```
./gradlew buildApplication
```

## Building for the current operating system
Execute the following on the command line:
```
./gradlew buildApplicationCurrentOperatingSystem
``` 