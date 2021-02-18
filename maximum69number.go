package main

import "strconv"

// 9966 9996  6999
func maximum69Number(num int) int {
	ans := 0
	countOfSix, curCountOfSix := 0, 0
	numCopy := num

	for numCopy > 0 {
		quotient, remainder := numCopy/10, numCopy%10
		if remainder == 6 {
			countOfSix++
		}
		numCopy = quotient
	}

	if countOfSix == 0 {
		return num
	}

	numCopy = num

	for numCopy > 0 {
		quotient, remainder := numCopy/10, numCopy%10
		if remainder == 6 {
			curCountOfSix++
			if curCountOfSix == countOfSix {
				ans = ans*10 + 9
				numCopy = quotient
				continue
			}
		}
		ans = ans*10 + remainder
		numCopy = quotient
	}

	ans, _ = strconv.Atoi(reverseString(strconv.Itoa(ans)))

	return ans
}
