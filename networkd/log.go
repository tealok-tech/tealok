package networkd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func extract_subnet(line string) string {
	re := regexp.MustCompile(`\b([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}/\d{1,3}\b|\b([0-9a-fA-F]{1,4}:){1,7}:/\d{1,3}\b|\b::([0-9a-fA-F]{1,4}:){1,6}[0-9a-fA-F]{1,4}/\d{1,3}\b|\b([0-9a-fA-F]{1,4}:){1,7}(:|:[0-9a-fA-F]{1,4}){0,6}/\d{1,3}\b`)
	match := re.FindString(line)
	// Regular expression to match an IPv6 CIDR notation
	return match
}
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
			return extract_subnet(lines[i])
		}
	}
	return ""
}
