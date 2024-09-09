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
	fmt.Println("[*] Starting integrity check ")

	// Ran CertUtil
	output_ct := utils.RanCertutil(filePath)

	fmt.Printf("Output: \n%s\n", output_ct)

	// Ran SignTool
	output_st := utils.RanSignTool(filePath)
	fmt.Printf("Output: \n%s\n", output_st)

}
