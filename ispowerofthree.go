package main

import (
	"regexp"
	"strconv"
)

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func isPowerOfThree2(n int) bool {
	if n < 1 {
		return false
	}

	num := int64(n)

	base3NumStr := strconv.FormatInt(num, 3)
	pattern, _ := regexp.Compile("^10*$")
	ans := pattern.MatchString(base3NumStr)

	return ans
}

func isPowerOfThree3(n int) bool {
	for i := 0; i <= 19; i++ {
		if n == power(3, i) {
			return true
		}
	}

	return false
}
