package ipv6

import (
	"testing"
)

func TestFormulateIpv6(t *testing.T) {
	basePrefix := "2001::/22"
	var bitWidth, maxSubPrefix = 6, 0
	ipv6Nets, err := FormulateIpv6(basePrefix, bitWidth, maxSubPrefix)
	if err != nil {
		t.Error(err)
	}

	t.Logf("basePrefix:%s total subNet:%d\n", basePrefix, len(ipv6Nets))
	for _, net := range ipv6Nets {
		t.Log(net.String())
	}
}
