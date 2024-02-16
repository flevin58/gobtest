package lib

import (
	"os/exec"
)

// Returns the serial number of the Windows PC
func GetComputerSerialNumber() string {
	out, _ := exec.Command("wmic", "bios", "get", "serialnumber").Output()
	return out
}
