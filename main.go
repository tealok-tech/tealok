package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("journalctl", "-u", "systemd-networkd", "-p", "6", "-g", "DHCP: received delegated prefix")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err);
		return
	}
	output := out.String()

	fmt.Println("Captured output: ")
	fmt.Println(output)
}
