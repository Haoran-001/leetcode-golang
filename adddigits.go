package main

import (
	"strconv"
)

func addDigits(num int) int {
	numStr := strconv.Itoa(num)

	for len(numStr) >= 2 {
		sum := 0
		for i := 0; i < len(numStr); i++ {
			sum += int(numStr[i]) - 48
		}
		numStr = strconv.Itoa(sum)
	}

	ans, _ := strconv.Atoi(numStr)
	return ans
}

func addDigits2(num int) int {
	ans := 0

	for {
		ans = 0
		for num > 0 {
			remainder, quotient := num%10, num/10
			ans += remainder
			num = quotient
		}

		if ans < 10 {
			return ans
		}

		num = ans
	}
}
