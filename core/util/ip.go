package util

import (
	"fmt"
	"strconv"
	"strings"
)

func GetIPType(ip string) string {
	if strings.Contains(ip, ".") {
		return getIPV4Type(ip)
	} else if strings.Contains(ip, ":") {
		return getIPV6Type(ip)
	} else {
		return ip
	}
}

func getIPV4Type(ip string) string {

	ipNums := strings.Split(ip, ".")
	if len(ipNums) != 4 {
		return ip
	}

	ipType := ""
	firstNum, err1 := strconv.Atoi(ipNums[0])
	secNum, err2 := strconv.Atoi(ipNums[1])
	thirdNum, err3 := strconv.Atoi(ipNums[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return ip
	}

	if firstNum > 0 && firstNum < 127 {
		ipType = fmt.Sprintf("ipv4-A-%d", firstNum)
	} else if firstNum > 127 && firstNum <= 191 {
		ipType = fmt.Sprintf("ipv4-B-%d-%d", firstNum, secNum)
	} else if firstNum > 191 && firstNum < 233 {
		ipType = fmt.Sprintf("ipv4-C-%d-%d-%d", firstNum, secNum, thirdNum)
	}

	return ipType
}

func getIPV6Type(ip string) string {

	ipNums := strings.Split(ip, ":")
	ipV6AllNums := []string{"0000", "0000", "0000", "0000",
		"0000", "0000", "0000", "0000"}
	suppleStrings := []string{"0000", "000", "00", "0", ""}

	i := 0
	for i < len(ipNums) {
		ipNum := ipNums[i]
		if len(ipNum) > 0 {
			ipNum = fmt.Sprintf("%s%s", suppleStrings[len(ipNum)], ipNum)
			ipV6AllNums[i] = ipNum
		} else {
			break
		}
		i++
	}

	j := len(ipNums) - 1
	numIndex := len(ipV6AllNums) - 1;
	for i < j {
		ipNum := ipNums[j]
		if len(ipNum) > 0 {
			ipNum = fmt.Sprintf("%s%s", suppleStrings[len(ipNum)], ipNum)
			ipV6AllNums[numIndex] = ipNum
		} else {
			break
		}
		j--
		numIndex--
	}

	return "ipv6-" + strings.Join(ipV6AllNums[0:2], "-")
}
