package netin

import (
	"bytes"
	"net"
)

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// check if a net.IP is an IPv4
func IsIP4(ip net.IP) bool {
	switch len(ip) {
	case net.IPv4len:
		return true
	case net.IPv6len:
		// check if IPv4-in-IPv6
		// This will return true for IPs formatted as: ::ffff:127.0.0.1
		return bytes.Equal(ip[:12], v4InV6Prefix)
	default:
		return false
	}
}
