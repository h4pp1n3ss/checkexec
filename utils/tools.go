package utils

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func FilterCommandOutput(command string, args []string, filter string) ([]string, error) {
	// Execute the command
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// Create a scanner to read the output line by line
	scanner := bufio.NewScanner(&out)
	var filteredLines []string

	// Loop through the lines and filter based on the given string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, filter) {
			filteredLines = append(filteredLines, line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filteredLines, nil
}
