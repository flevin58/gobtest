package lib

import (
	"os/exec"
	"strings"
)

// Linux: dmidecode -t system
// Returns the serial number of the Mac
func GetComputerSerialNumber() string {
	out, err := exec.Command("/bin/bash", "-c", "ioreg -l | grep IOPlatformSerialNumber").Output()
	if err != nil {
		return ""
	}
	_, after, found := strings.Cut(string(out), "=")
	if !found {
		return ""
	}
	result := strings.TrimSpace(after)
	return strings.Replace(result, `"`, ``, 2)
}
