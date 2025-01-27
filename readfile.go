package main

import (
	"bufio"
	"os"
	"strings"
)

// ReadFile reads the given file and returns its contents as a list of lines
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var orgLines []string
	for scanner.Scan() {
		orgLines = append(orgLines, strings.Replace(strings.TrimSpace(scanner.Text()), "/r/n", "/n", -1))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return orgLines, nil
}
