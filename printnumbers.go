package main

func printNumbers(n int) []int {
	endNumber := 9
	n = n - 1
	for n > 0 {
		endNumber = endNumber*10 + 9
		n--
	}
	ans := make([]int, endNumber)

	for i := 1; i <= endNumber; i++ {
		ans[i-1] = i
	}

	return ans
}
