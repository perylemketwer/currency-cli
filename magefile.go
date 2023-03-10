//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// Constant to OS types to compile
var listOs = []string{
	"darwin",
	"dragonfly",
	"freebsd",
	"linux",
	"netbsd",
	"openbsd",
	"solaris",
	"windows",
}

// Create a local binary to execute
func Build() error {
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "./bin/currency", ".")
	return cmd.Run()
}

// Compiling a binary file to OS with amd64, arm64 or 386 architecture:
// e.g mage compile linux amd64
func Compile(system, arch string) error {
	fmt.Printf("Compliling to %s to %s architecture...\n", system, arch)

	var binName string
	if system == "windows" {
		binName = fmt.Sprintf("./bin/currency-%s-%s/currency.exe", system, arch)
	} else {
		binName = fmt.Sprintf("./bin/currency-%s-%s/currency", system, arch)
	}

	systemName := fmt.Sprintf("GOOS=%s", system)
	archName := fmt.Sprintf("GOARCH=%s", arch)

	cmd := exec.Command("go", "build", "-o", binName, ".")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, systemName)
	cmd.Env = append(cmd.Env, archName)

	return cmd.Run()
}

// Compiling a binary to cross platform: Linux, MacOS, Windows, OpenBSD and FreeBSD
func CompileAll() {
	fmt.Println("Compliling for every OS and Platform...")

	for _, k := range listOs {
		switch k {
		case "darwin":
			Compile(k, "arm64")
			Compile(k, "amd64")
		case "dragonfly":
			Compile(k, "amd64")
		default:
			Compile(k, "amd64")
			Compile(k, "386")
			Compile(k, "arm64")
		}
	}

	fmt.Println("Done!")
}

// Running a CLI with currency type
func Run(coin string) {
	output, err := exec.Command("go", "run", "main.go", coin).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

// Running unit tests
func Tests() {
	fmt.Println("Running tests...")
	output, err := exec.Command("go", "test", "-v", "./tests").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

// Remove all files of the project
func Clean() {
	fmt.Println("Cleaning...")

	err := os.RemoveAll("./bin")
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll("./dist")
	if err != nil {
		log.Fatal(err)
	}
}
