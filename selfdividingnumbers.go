package main

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			ans = append(ans, i)
		}
	}

	return ans
}

func isSelfDividingNumber(num int) bool {
	numCopy := num
	for numCopy > 0 {
		remainder, quotient := numCopy%10, numCopy/10
		if remainder == 0 || num%remainder != 0 {
			return false
		}
		numCopy = quotient
	}

	return true
}
