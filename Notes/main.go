package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command) // Use bash to execute the command
	output, err := cmd.CombinedOutput()        // Capture standard output and error
	if err != nil {
		return string(output), fmt.Errorf("error executing command: %v", err)
	}
	return string(output), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command to execute: ")
	command, err := reader.ReadString('\n') // Read input until newline
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	command = command[:len(command)-1]

	output, err := ExecuteCommand(command)
	if err != nil {
		fmt.Println("Command execution failed:", err)
		fmt.Println("Error output:", output)
		return
	}

	fmt.Println("Command output:")
	fmt.Println(output)
}
