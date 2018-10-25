package validate

import (
	"net"
)

func ValidateIP(ip string) bool {
	bs := net.ParseIP(ip)
	return !(bs == nil)
}
