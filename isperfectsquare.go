package main

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	if num == 1 {
		return true
	}

	low, high := 1, num/2

	for low <= high {
		mid := (low + high) / 2
		guessSquared := mid * mid
		if guessSquared == num {
			return true
		} else if guessSquared < num {
			low = mid + 1
		} else if guessSquared > num {
			high = mid - 1
		}
	}

	return false
}
