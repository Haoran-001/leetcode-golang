package main

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return n&-n == n
}

func isPowerOfTwo2(n int) bool {
	//powerOfTwoArray := make([]int, 32)

	for i := 0; i <= 31; i++ {
		if 1<<i == n {
			return true
		}
	}

	return false
}

func isPowerOfTwo3(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}
