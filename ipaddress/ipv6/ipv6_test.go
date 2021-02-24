package ipv6

import (
	"testing"
)

func TestFormulateIpv6(t *testing.T) {
	basePrefix := "2002::/32"
	var bitWidth, maxSubPrefix = 12, 0
	ipv6Nets, err := FormulateIpv6(basePrefix, bitWidth, maxSubPrefix)
	if err != nil {
		t.Error(err)
	}

	t.Logf("basePrefix:%s total subNet:%d\n", basePrefix, len(ipv6Nets))
	for i, net := range ipv6Nets {
		t.Logf("index:%d %s\n", i, net.String())
	}

	//b1 := 0xfc00
	//b2 := 0xfc00
	//fmt.Printf("b1:%b b2:%b (b1&b2):%b (b2&b1==b1):%t \n", b1, b2, b2&b1, b2&b1 == b1)

	prefixBeginValue := "f0"
	if contain, err := CheckValueInPrefix(basePrefix, bitWidth, prefixBeginValue); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("contain:%t", contain)
	}
}
