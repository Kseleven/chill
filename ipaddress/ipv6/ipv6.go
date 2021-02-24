package ipv6

import (
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"strings"
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

func CheckValueInPrefix(basePrefix string, bitWidth int, prefixBeginValue string) (bool, error) {
	baseIp, baseIpNet, err := net.ParseCIDR(basePrefix)
	if err != nil {
		return false, err
	}
	ones, size := baseIpNet.Mask.Size()
	subPrefixLength := ones + bitWidth

	if subPrefixLength > MaxIpv6FormulateLength {
		return false, fmt.Errorf("prefix length should smaller than %d", MaxIpv6FormulateLength)
	}

	ipv6Integer := ipv6ToInteger(baseIp)
	subInteger, err := parsePrefixBeginValue(prefixBeginValue, bitWidth, subPrefixLength)
	if err != nil {
		return false, err
	}
	reverseIp := integerToIpv6(ipv6Integer | subInteger<<(MaxIpv6FormulateLength-subPrefixLength))
	tempIpNet := net.IPNet{IP: reverseIp, Mask: net.CIDRMask(subPrefixLength, size)}

	fmt.Println("subInteger", subInteger)
	return baseIpNet.Contains(tempIpNet.IP), nil
}

func parsePrefixBeginValue(beginValue string, bitWidth, subPrefixLength int) (uint64, error) {
	total := calculateLeftPrefixCount(0, bitWidth, 1)
	beginValueSlice := make([]byte, 0)
	for i := 0; i < len(beginValue); i++ {
		v, err := hex.DecodeString(fromHexChar(beginValue[i]))
		if err != nil {
			return 0, err
		}
		beginValueSlice = append(beginValueSlice, v[0])
	}

	bytesInteger := bytesToInteger(beginValueSlice)
	cardinalValue := (len(trimSuffixZero(beginValue))/4 + 1) * 16
	rightPoint := cardinalValue - subPrefixLength%cardinalValue
	maxValue := 0 | uint64(total)<<rightPoint
	if maxValue < bytesInteger {
		return 0, fmt.Errorf("prefix %s value is too bigger than %x", beginValue, maxValue)
	}

	if (maxValue & bytesInteger) == bytesInteger {
		return bytesInteger >> rightPoint, nil
	}

	return 0, fmt.Errorf("not found prefix value %s", beginValue)
}

func bytesToInteger(slice []byte) uint64 {
	var out uint64
	for i := 0; i < len(slice); i++ {
		out |= uint64(slice[i]) << (4 * (len(slice) - 1 - i))
	}
	return out
}

func fromHexChar(c byte) string {
	return string([]byte{'0', c})
}

func trimSuffixZero(prefixBeginValue string) string {
	value := strings.TrimSuffix(prefixBeginValue, "0")
	for value != strings.TrimSuffix(value, "0") {
		value = strings.TrimSuffix(value, "0")
	}
	return value
}

func calculateLeftPrefixCount(usedNumber, bitWidth, prefixNumber int) int64 {
	return int64(int(math.Pow(2, float64(bitWidth)))*prefixNumber - prefixNumber - usedNumber)
}
