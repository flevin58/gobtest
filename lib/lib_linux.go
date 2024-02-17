package lib

import (
	"os/exec"
	"strings"
)

// Returns the serial number of the linux computer
func GetComputerSerialNumber() string {
	out, err := exec.Command("/usr/bin/sudo", "bash", "-c", "dmidecode -t system | grep \"Serial Number:\"").Output()
	if err != nil {
		return ""
	}
	outs := string(out)
	index := strings.Index(outs, ":")
	return strings.TrimSpace(outs[index+1:])
}
