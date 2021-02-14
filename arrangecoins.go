package main

func arrangeCoins(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if n-sum < 0 {
			return i - 1
		}
	}

	return -1
}

// 1 + 2 + 3 + ... + n = n * (n + 1) / 2    5
func arrangeCoins2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	for i := 0; i <= n; i++ {
		if (i*(i+1))/2 > n {
			return i - 1
		}
	}

	return -1
}

// 5 1-5 3 1-2 1 2-2 2 3-2
func arrangeCoins3(n int) int {
	low, high := 1, n

	for low <= high {
		mid := (low + high) / 2
		sum := (mid * (mid + 1)) / 2

		if sum == n {
			return mid
		} else if sum < n {
			low = mid + 1
		} else if sum > n {
			high = mid - 1
		}
	}

	return high
}
