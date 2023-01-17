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

	binName := fmt.Sprintf("./bin/currency-%s-%s", system, arch)
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

	listOs := []string{"darwin", "freebsd", "linux", "openbsd", "windows"}
	for _, k := range listOs {
		Compile(k, "amd64")
		if k == "darwin" {
			Compile(k, "arm64")
			continue
		}
		Compile(k, "386")
	}

	fmt.Println("Done!")
}

// Running a CLI with currency type
func Run(coin string) {
	fmt.Println("Running...")
	output, err := exec.Command("go", "run", "main.go", coin).Output()
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
}
