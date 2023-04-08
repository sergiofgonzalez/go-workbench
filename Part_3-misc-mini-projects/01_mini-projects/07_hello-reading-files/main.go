package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Using `os.ReadFile`
	lines, err := readLines("./data/data.txt")
	if err != nil {
		fmt.Printf("Unable to read lines from file: %v\n", err)
		return
	}
	for _, line := range lines {
		fmt.Println(line)
	}

	// Using `os.Open()` followed by `f.Read()`
	text, err := readBytes("./data/data.txt")
	if err != nil {
		fmt.Printf("Unable to read contents from file: %v\n", err)
		return
	}
	fmt.Println(text)

	// Reading line by line
	f, err := os.Open("./data/data.txt")
	if err != nil {
		fmt.Println("Unable to open file")
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, err := readLineByLine(reader)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Printf("Unable to read line by line: %v\n", err)
				return
			} else {
				break
			}
		} else {
			fmt.Println(line)
		}
	}

	// Reading using scanners, which is more high-level
	contents, err := readByScanning("./data/data.txt")
	if err != nil {
		fmt.Printf("Unable to read contents from file: %v\n", err)
		return
	}
	fmt.Println(contents)
}

// readLines read the whole file in one shot, materializing the whole file in memory
// which is not a good idea if the file is really large
func readLines(pathToFile string) ([]string, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read file %q: %w", pathToFile, err)
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

// readBytes is an alternative implementation that does not materialize the whole file
// but it is even more "low-level" than `readLines`
func readBytes(pathToFile string) (string, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return "", fmt.Errorf("unable to open %q: %w", pathToFile, err)
	}
	defer f.Close()

	var data []byte
	buf := make([]byte, 16)	// read in chunks of 16 bytes
	for {
		fmt.Printf("Reading chunk: ")
		n, err := f.Read(buf)
		fmt.Printf("OK\n")
		chunk := buf[0:n]
		data = append(data, chunk...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("unable to read %q: %w", pathToFile, err)
		}
	}
	return string(data), nil
}

func readLineByLine(reader *bufio.Reader) (string, error) {
	var lineBytes []byte
	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			return "", fmt.Errorf("could not read line from reader: %w", err)
		}
		lineBytes = append(lineBytes, data...)
		if !isPrefix {
			break
		}
	}

	line := string(lineBytes)
	return line, nil
}

func readByScanning(pathToFile string) (string, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return "", fmt.Errorf("unable to open %q: %w", pathToFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sb := strings.Builder{}
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}
	return sb.String(), nil
}