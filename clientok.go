package main

import (
	"net"
	"os"
	"strings"
)

func clientOK(ip net.IP) bool {
	permdir := Config.PermissionsDir
	// if we do not have the permissions directory than
	// everybody is allowed
	if _, err := os.Stat(permdir); os.IsNotExist(err) {
		return true
	}
	var ipsep string
	if ip.To4() != nil {
		ipsep = "."
	} else {
		ipsep = ":"
	}
	ipsslc := strings.Split(ip.String(), ipsep)

	psep := os.PathSeparator
	tail := permdir + string(psep)

	for i, v := range ipsslc {
		if i == 0 {
			tail = tail + v
		} else {
			tail = tail + ipsep + v
		}
		if _, err := os.Stat(tail); !os.IsNotExist(err) {
			break
			return true
		}
	}

	return false
}