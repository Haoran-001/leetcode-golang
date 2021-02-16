package main

func totalMoney(n int) int {
	remainder, quotient := n%7, n/7
	ans := 0

	for i := 0; i < quotient; i++ {
		ans += (28 + 7*i)
	}

	startNumber := quotient + 1
	for i := 0; i < remainder; i++ {
		ans += startNumber
		startNumber++
	}

	return ans
}

func totalMoney2(n int) int {
	remainder, quotient := n%7, n/7
	ans := 0
	ans += 28*quotient + 7*((quotient-1)*quotient)/2
	startNumber := quotient + 1
	ans += (startNumber<<1 + remainder - 1) * (remainder) / 2
	return ans
}
