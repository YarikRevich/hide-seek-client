package creators

import (
	"strings"
)

func CreateAddr(ip string, port string) (string, error) {
	var addr strings.Builder
	for i := 0; i < len(ip); i++ {
		if err := addr.WriteByte(ip[i]); err != nil {
			return "", err
		}
	}
	addr.WriteString(":")

	for i := 0; i < len(port); i++ {
		if err := addr.WriteByte(port[i]); err != nil {
			return "", err
		}
	}

	return addr.String(), nil
}
