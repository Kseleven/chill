package ipv6

import (
	"fmt"
	"math"
	"net"
)

const (
	MaxIpv6FormulateLength = 64
)

func FormulateIpv6(basePrefix string, bitWidth, maxSubPrefix int) ([]net.IPNet, error) {
	ip, baseIpNet, err := net.ParseCIDR(basePrefix)
	if err != nil {
		return nil, err
	}
	ipv6Integer := ipv6ToInteger(ip)
	ones, size := baseIpNet.Mask.Size()
	subPrefixLength := ones + bitWidth

	if subPrefixLength > MaxIpv6FormulateLength {
		return nil, fmt.Errorf("prefix length should smaller than %d", MaxIpv6FormulateLength)
	}

	if totalSubPrefix := int(math.Pow(2, float64(bitWidth))); maxSubPrefix > totalSubPrefix {
		return nil, fmt.Errorf("subPrefix numbers %d is bigger than %d", maxSubPrefix, totalSubPrefix)
	} else if maxSubPrefix == 0 {
		maxSubPrefix = totalSubPrefix
	}

	var ipNets []net.IPNet
	for i := 0; i < maxSubPrefix; i++ {
		temp := ipv6Integer
		temp |= uint64(i) << (MaxIpv6FormulateLength - subPrefixLength)
		reverseIp := integerToIpv6(temp)
		ipNets = append(ipNets, net.IPNet{IP: reverseIp, Mask: net.CIDRMask(subPrefixLength, size)})
	}

	return ipNets, nil
}

func ipv6ToInteger(ip net.IP) uint64 {
	return uint64(ip[0])<<56 | uint64(ip[1])<<48 | uint64(ip[2])<<40 | uint64(ip[3])<<32 |
		uint64(ip[4])<<24 | uint64(ip[5])<<16 | uint64(ip[6])<<8 | uint64(ip[7])
}

func integerToIpv6(value uint64) net.IP {
	ip := make(net.IP, net.IPv6len)
	ip[0] = byte(value >> 56)
	ip[1] = byte(value >> 48)
	ip[2] = byte(value >> 40)
	ip[3] = byte(value >> 32)
	ip[4] = byte(value >> 24)
	ip[5] = byte(value >> 16)
	ip[6] = byte(value >> 8)
	ip[7] = byte(value)
	return ip
}
