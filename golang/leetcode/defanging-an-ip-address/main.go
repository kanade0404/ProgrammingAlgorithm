package defanging_an_ip_address

import "strings"

func defangIPaddr(address string) string {
	strs := strings.Split(address, "")
	var result string
	for _, str := range strs {
		if str == "." {
			result = result + "[.]"
		} else {
			result = result + str
		}
	}
	return result
}
