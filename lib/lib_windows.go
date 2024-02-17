package lib

import (
	"os/exec"
	"strings"
)

// Returns the serial number of the Windows PC
func GetComputerSerialNumber() string {
	out, err := exec.Command("wmic", "bios", "get", "serialnumber").Output()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(out), "\r\n")
	if strings.TrimSpace(lines[0]) != "SerialNumber" {
		return ""
	}
	return strings.TrimSpace(lines[1])
}
