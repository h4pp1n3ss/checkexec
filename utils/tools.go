package utils

import (
	"log"
	"os/exec"
)

func RanCertutil(filePath string) string {

	hashAlgorithm := "SHA256"

	// Construct the certutil command
	cmd := exec.Command("certutil", "-hashfile", filePath, hashAlgorithm)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run certutil: %v", err)
	}
	// Return output
	return string(output)

}

func RanSignTool(filePath string) string {

	cmd := exec.Command("signtool", "verify", "/pa", "/v", filePath)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run signtool: %v", err)
	}

	// Return output
	return string(output)

}
