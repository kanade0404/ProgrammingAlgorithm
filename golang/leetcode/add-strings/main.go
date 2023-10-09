package add_strings

import (
	"log"
	"math"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l := int(math.Max(float64(len(num1)), float64(len(num2))))
	var (
		result   string
		overflow bool
	)
	for i := 0; i < l; i++ {
		var tmp int
		if i < len(num1) {
			if n1, err := strconv.Atoi(string(rune(num1[len(num1)-i-1]))); err != nil {
				log.Fatalln(err)
			} else {
				tmp += n1
			}
		}
		if i < len(num2) {
			if n2, err := strconv.Atoi(string(rune(num2[len(num2)-i-1]))); err != nil {
				log.Fatalln(err)
			} else {
				tmp += n2
			}
		}
		if overflow {
			tmp += 1
		}
		if tmp >= 10 {
			overflow = true
			result = strconv.Itoa(tmp%10) + result
		} else {
			overflow = false
			result = strconv.Itoa(tmp) + result
		}
	}
	if overflow {
		result = "1" + result
	}
	return result
}
