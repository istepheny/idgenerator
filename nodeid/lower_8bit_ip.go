package nodeid

import (
	"errors"
	"net"
)

func Lower8BitIP() (uint8, error) {
	ip, err := getIP()
	if err != nil {
		return 0, err
	}

	return ip[3], nil
}

func getIP() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() || ipnet.IP.To4() == nil {
			continue
		}

		return ipnet.IP.To4(), nil
	}

	return nil, errors.New("no ip address")
}
