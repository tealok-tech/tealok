package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	
	cmd := exec.Command("journalctl", "-u", "systemd-networkd", "-p", "6", "-g", "DHCP: received delegated prefix")

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := exitError.ExitCode()
			fmt.Println("Exit code:", exitCode)
			os.Exit(exitCode)
		} else {
			fmt.Println("Command failed with:", err)
			fmt.Println("Captured stderr:", stderr.String())
			fmt.Println("Captured stdout:", out.String())
			os.Exit(100)
		}
	}
	output := out.String()

	fmt.Println("Captured output: ")
	fmt.Println(output)
}
