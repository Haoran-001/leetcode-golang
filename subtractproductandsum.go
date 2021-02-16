package main

import (
	"strconv"
)

func subtractProductAndSum(n int) int {
	nStr := strconv.Itoa(n)
	multiplication := 1
	addition := 0

	for _, item := range nStr {
		multiplication *= int(item) - 48
		addition += int(item) - 48
	}

	ans := multiplication - addition
	return ans
}

func subtractProductAndSum2(n int) int {
	var multiplication int = 1
	var addition int = 0

	for n > 0 {
		remainder := n % 10
		multiplication *= remainder
		addition += remainder
		quotient := n / 10
		n = quotient
	}

	var ans int = multiplication - addition

	return ans
}
