// +build !linux,!windows

package netin

import (
	"errors"
	"syscall"
)

// fallback for unsupported OSes
const (
	FamilyIPv4 = 0x0800
	FamilyIPv6 = 0x86dd
)

// GetFamily returns the family for a given socket. It can be then compared
// with the FamilyIPv4 or FamilyIPv6 constants.
func GetFamily(c syscall.Conn) (uint16, error) {
	var err2 error
	var family uint16

	sc, err := c.SyscallConn()
	if err != nil {
		return family, err
	}

	err = sc.Control(func(fd uintptr) {
		var sa syscall.Sockaddr

		sa, err2 = syscall.Getsockname(int(fd))
		if err2 != nil {
			return
		}

		switch sa.(type) {
		case *syscall.SockaddrInet4:
			family = FamilyIPv4
		case *syscall.SockaddrInet6:
			family = FamilyIPv6
		default:
			err2 = errors.New("unknown address family")
		}
	})

	if err2 != nil {
		return family, err2
	}
	return family, err
}
