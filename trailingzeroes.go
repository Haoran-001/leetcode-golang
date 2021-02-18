package main

func trailingZeroes(n int) int {
	var ansMap map[int]int = make(map[int]int, 2)

	for i := 5; i <= n; i *= 5 {
		ansMap[5] += n / i
	}

	for i := 2; i <= n; i *= 2 {
		ansMap[2] += n / 2
	}

	return min(ansMap[2], ansMap[5])
}

func trailingZeroes2(n int) int {
	var ans int = 0

	for i := 5; i <= n; i *= 5 {
		ans += n / i
	}

	return ans
}
