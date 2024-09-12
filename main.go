package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
)

func printBanner(){
	fmt.Println("  ______  __    __   _______   ______  __  ___  __________   ___  _______   ______ ")
	fmt.Println(" /      ||  |  |  | |   ____| /      ||  |/  / |   ____\\  \\ /  / |   ____| /      |")
	fmt.Println("|  ,----'|  |__|  | |  |__   |  ,----'|  '  /  |  |__   \\  V  /  |  |__   |  ,----'")
	fmt.Println("|  |     |   __   | |   __|  |  |     |    <   |   __|   >   <   |   __|  |  |     ")
	fmt.Println("|  `----.|  |  |  | |  |____ |  `----.|  .  \\  |  |____ /  .  \\  |  |____ |  `----.")
	fmt.Println(" \\______||__|  |__| |_______| \\______||__|\\__\\ |_______/__/ \\__\\ |_______| \\______|")
	fmt.Println("                                                                                  ")
	fmt.Println("Description: Simple wrapper of certutil and signtool for Hash and Signature Verification.")
	fmt.Println("Version: v0.1")
	fmt.Println("\n")
}

// Function to execute the certutil command and get the SHA256 hash
func executeCertutilCommand(filePath string) (string, error) {
	cmd := exec.Command("certutil", "-hashfile", filePath, "SHA256")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Function to extract and print the SHA256 hash from certutil output
func printSHA256Hash(output string) {
	// certutil outputs multiple lines, and the hash is in the second line
	lines := strings.Split(output, "\n")
	if len(lines) >= 2 {
		// The second line should contain the hash
		hash := strings.TrimSpace(lines[1])
		fmt.Println("[+] SHA256 Hash:", hash)
	} else {
		fmt.Println("Error: Could not extract the SHA256 hash.")
	}
}

// Function to execute the signtool verify command
func executeSigntoolCommand(filePath string) (string, error) {
	cmd := exec.Command("signtool", "verify", "/pa", "/v", filePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Function to save both the hash and signtool output to a file
func saveOutputToFile(hash string, signtoolOutput string, outputFilePath string) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("[!] ================== SHA256 Hash:\n")
	if err != nil {
		return err
	}
	_, err = writer.WriteString(hash + "\n\n")
	if err != nil {
		return err
	}

	_, err = writer.WriteString("[!] ==================  Signtool Verification Output:\n")
	if err != nil {
		return err
	}
	_, err = writer.WriteString(signtoolOutput + "\n")
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

// Function to get file path from command-line arguments
func getFilePathFromArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("file path not provided")
	}
	return os.Args[1], nil
}

// Function to read the output file and check for signtool verification result
func readAndVerifyOutput(outputFilePath string) error {
	content, err := ioutil.ReadFile(outputFilePath)
	if err != nil {
		return err
	}

	output := string(content)

	// Check if the string "Number of files successfully Verified: 1" is present
	if strings.Contains(output, "Number of files successfully Verified: 1") {
		fmt.Println("[+] Signtool verification successful.")
	} else {
		fmt.Println("[-] There are some warnings in the file, please check manually.")
	}

	return nil
}

// Main function
func main() {
	// Step 0: Print banner
	printBanner()

	// Step 1: Get the file path from command-line arguments
	filePath, err := getFilePathFromArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: Execute certutil command to get the SHA256 hash
	hashOutput, err := executeCertutilCommand(filePath)
	if err != nil {
		fmt.Println("Error executing certutil command:", err)
		return
	}

	// Step 3: Print only the SHA256 hash
	printSHA256Hash(hashOutput)

	// Step 4: Execute signtool command to verify the file's signature
	signtoolOutput, err := executeSigntoolCommand(filePath)
	if err != nil {
		fmt.Println("Error executing signtool command:", err)
		return
	}

	// Step 5: Save both the hash and signtool output to a file
	outputFilePath := "output.txt"
	err = saveOutputToFile(hashOutput, signtoolOutput, outputFilePath)
	if err != nil {
		fmt.Println("Error saving output to file:", err)
		return
	}

	fmt.Printf("[+] SHA256 hash and signtool verification output saved to %s\n", outputFilePath)

	// Step 6: Read and verify the output file for signtool verification result
	err = readAndVerifyOutput(outputFilePath)
	if err != nil {
		fmt.Println("Error reading and verifying output file:", err)
	}
}
