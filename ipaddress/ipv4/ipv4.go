package ipv4

import (
	"fmt"
	"math"
	"net"
)

const (
	MaxIpv4FormulateLength = 32
)

func FormulateIpv4(basePrefix string, bitWidth, maxSubPrefix int) ([]net.IPNet, error) {
	ip, baseIpNet, err := net.ParseCIDR(basePrefix)
	if err != nil {
		return nil, err
	}
	ipv4Value := toIpv4Value(ip)
	ones, size := baseIpNet.Mask.Size()
	subPrefixLength := ones + bitWidth

	if subPrefixLength > MaxIpv4FormulateLength {
		return nil, fmt.Errorf("prefix length should smaller than %d", MaxIpv4FormulateLength)
	}

	if totalSubPrefix := int(math.Pow(2, float64(bitWidth))); maxSubPrefix > totalSubPrefix {
		return nil, fmt.Errorf("subPrefix numbers %d is bigger than %d", maxSubPrefix, totalSubPrefix)
	} else if maxSubPrefix == 0 {
		maxSubPrefix = totalSubPrefix
	}

	var ipNets []net.IPNet
	for i := 0; i < maxSubPrefix; i++ {
		temp := ipv4Value
		temp |= uint32(i) << (MaxIpv4FormulateLength - subPrefixLength)
		reverseIp := ipv4ValueToIp(temp)
		ipNets = append(ipNets, net.IPNet{IP: reverseIp, Mask: net.CIDRMask(subPrefixLength, size)})
	}

	return ipNets, nil
}

func toIpv4Value(ip net.IP) uint32 {
	ipv4 := ip.To4()
	return uint32(ipv4[0])<<24 | uint32(ipv4[1])<<16 | uint32(ipv4[2])<<8 | uint32(ipv4[3])
}

func ipv4ValueToIp(value uint32) net.IP {
	return net.IPv4(byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}
