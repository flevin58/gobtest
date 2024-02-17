package lib

import (
	"strings"
)

func GetSecretKey() [32]byte {
	// We're using a 32 byte long secret key
	var result [32]byte
	sn := GetComputerSerialNumber()
	count := (32 / len(sn)) + 1
	repeated := strings.Repeat(sn, count)
	copy(result[:], repeated[:32])
	return result
}
