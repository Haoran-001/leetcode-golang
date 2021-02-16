package main

func smallestRangeI(A []int, K int) int {
	minValue, maxValue := getMinArrayElements(A), getMaxArrayElements(A)
	ans := 0

	if maxValue-minValue >= 2*K {
		ans = maxValue - minValue - 2*K
	}

	return ans
}
