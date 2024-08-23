package networkd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Subnet() string {
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

	lines := strings.Split(output, "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		if strings.Contains(lines[i], "delegated prefix") {
			fmt.Println("Matching line:", lines[i])
			return lines[i]
		}
	}
	return ""
}
