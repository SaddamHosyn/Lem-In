package main

import (
	"bufio"
	"os"
)

// ReadFile reads the given file and returns its contents as a list of lines
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var originalFileLines []string
	for scanner.Scan() {
		originalFileLines = append(originalFileLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return originalFileLines, nil
}
