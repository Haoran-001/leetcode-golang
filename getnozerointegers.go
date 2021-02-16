package main

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		a := strconv.Itoa(i)
		b := strconv.Itoa(n - i)
		if !strings.Contains(a, "0") && !strings.Contains(b, "0") {
			return []int{i, n - i}
		}
	}

	return []int{0, 0}
}
