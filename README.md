# CheckExec
## Hash and Signature Verification Tool

This Go program performs the following tasks:
1. Calculates the SHA256 hash of a specified file using the `certutil` command.
2. Verifies the signature of the file using the `signtool verify /pa /v` command.
3. Saves both the SHA256 hash and the `signtool` verification output to a file.
4. Displays the SHA256 hash in the console.
5. Checks if the file was successfully verified by the `signtool`. If not, a warning message is displayed in the console.

## Requirements

- **Go**: The program is written in Go, so you need Go installed on your machine.
- **certutil**: The program uses the `certutil` command to calculate the SHA256 hash. Make sure it is available in your system's PATH.
- **signtool**: The program uses the `signtool` command to verify the file signature. Make sure it is installed and available in your system's PATH.

## Installation

1. Install Go on your system: [Go installation guide](https://golang.org/doc/install).
2. Clone this repository or download the `hashfile.go` file.
3. Make sure `certutil` and `signtool` are available in your system's PATH.

## Usage

To use the program, simply run it from the command line with the file path you want to check.

### Command

```bash
go run checkexec.go <filePath>
