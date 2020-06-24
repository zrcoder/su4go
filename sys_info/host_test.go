package sys_info

import "testing"

func TestExternalIP(t *testing.T) {
	t.Log(ExternalIPOfType(IpV4))
	t.Log(ExternalIPOfType(IpV6))
	t.Log(ExternalIPOfType(3))
}
