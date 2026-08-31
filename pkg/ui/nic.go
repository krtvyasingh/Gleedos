package ui

import "net"

func ListInterfaces() int {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0
	}
	return len(ifaces)
}
