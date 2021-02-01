package ipv4

import (
	"testing"
)

func TestFormulateIpv4(t *testing.T) {
	basePrefix := "10.0.0.0/16"
	var bitWidth, maxSubPrefix = 8, 0
	ipv4Nets, err := FormulateIpv4(basePrefix, bitWidth, maxSubPrefix)
	if err != nil {
		t.Error(err)
	}

	t.Logf("basePrefix:%s total subNet:%d\n", basePrefix, len(ipv4Nets))
	for _, net := range ipv4Nets {
		t.Log(net.String())
	}
}
