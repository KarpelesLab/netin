package netin

import (
	"syscall"
	"unsafe"
)

const (
	FamilyIPv4 = syscall.AF_INET
	FamilyIPv6 = syscall.AF_INET6
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
		var sa syscall.RawSockaddrAny
		var l uint // type Socklen uint

		// int getsockname(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
		_, _, err2 = syscall.Syscall(syscall.SYS_GETSOCKNAME, fd, uintptr(unsafe.Pointer(&sa)), uintptr(unsafe.Pointer(&l)))
		if err2 != nil {
			return
		}

		family = sa.Addr.Family
	})

	if err2 != nil {
		return family, err2
	}
	return family, err
}
