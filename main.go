package main

import (
	"fmt"
	"log"
	"os"
	"secwrapper/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file path as an argument.")
	}

	filePath := os.Args[1]

	// Step 1. Verify file integrity using CertUtil Tool
	fmt.Println("[*] Executing certutil")

	// "certutil", "-hashfile", filePath, hashAlgorithm
	filteredOutputCertutil, err := utils.FilterCommandOutput("certutil", []string{"-hashfile", filePath, "SHA256"}, "valid")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the filtered output
	for _, line := range filteredOutputCertutil {
		fmt.Println(line)
	}

	// cmd := exec.Command("signtool", "verify", "/pa", "/v", filePath)
	fmt.Println("[*] Executing signtool")
	filteredOutputSigntool, err := utils.FilterCommandOutput("signtool", []string{"verify", "/pa", "/v", filePath}, "valid")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the filtered output
	for _, line := range filteredOutputSigntool {
		fmt.Println(line)
	}

}
