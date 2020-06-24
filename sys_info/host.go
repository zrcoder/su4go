package sys_info

import (
	"net"
	"strings"
)

type IpType int

const (
	IpV4 IpType = 0
	IpV6 IpType = 1
)

func ExternalIP() string {
	return ExternalIPOfType(IpV4)
}

//get host ip address, if an error occured, returns ""
func ExternalIPOfType(ipType IpType) string {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 {
			continue
		}
		if i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, _ := i.Addrs()
		for _, addr := range addresses {
			var ip net.IP = nil
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			if ipType == IpV4 && ip.To4() != nil {
				return ip.String()
			}
			if ipType == IpV6 && strings.Contains(ip.String(), ":") {
				return ip.String()
			}
		}
	}
	return ""
}
