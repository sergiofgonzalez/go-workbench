package main

import (
	"errors"
	"fmt"
)

type errFileNotFound struct {
	filename string
}

func (err errFileNotFound) Error() string {
	return fmt.Sprintf("file %q was not found", err.filename)
}

type errNoSuchFile string

func (err errNoSuchFile) Error() string {
	return string(err)
}

// ErrorNoDataFile is a sentinel error when data.txt is not found
const ErrorNoDataFile = errNoSuchFile("No such file data.txt")

func main() {

	// Non-sentinel error
	fmt.Println("== Non-sentinel")
	err := triggerFileNotFoundError()
	if err != nil {
		fmt.Println("An error was found, but it cannot be used as a sentinel error")
		// if errors.Is(err, errFileNotFound) {}  // Doesn't compile
	}
	fmt.Println()

	// wrapped sentinel error
	fmt.Println("== Sentinel error wrapped")
	err = triggerWrappedErrorNoDataFile()
	if err != nil {
		fmt.Println("An error was found...")
		if err == ErrorNoDataFile {
			fmt.Println("The error was exactly an ErrorNoDataFile")
		}
		if errors.Is(err, ErrorNoDataFile) {
			fmt.Println("The error wrapped an ErrorNoDataFile")
		}
	}
	fmt.Println()

	// sentinel error
	fmt.Println("== Sentinel error")
	err = triggerErrorNoDataFile()
	if err != nil {
		fmt.Println("An error was found...")
		if err == ErrorNoDataFile {
			fmt.Println("The error was exactly an ErrorNoDataFile")
		}
		if errors.Is(err, ErrorNoDataFile) {
			fmt.Println("The error wrapped an ErrorNoDataFile")
		}
	}
	fmt.Println()
}

func triggerFileNotFoundError() error {
	return fmt.Errorf("triggering error %w", errFileNotFound{"data.txt"})
}

func triggerWrappedErrorNoDataFile() error {
	return fmt.Errorf("triggering error %w", ErrorNoDataFile)
}

func triggerErrorNoDataFile() error {
	return ErrorNoDataFile
}