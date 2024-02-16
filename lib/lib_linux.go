package lib

import (
	"os/exec"
)

// Returns the serial number of the linux computer
func GetComputerSerialNumber() string {
	out, _ := exec.Command("dmidecode", "-t", "system").Output()
	return out
}
