package main

import "math"

// x^(1/2) = e^(1/2*lnx)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	ans := int(math.Exp(0.5 * math.Log(float64(x))))

	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}

	return ans
}

func mySqrt2(x int) int {
	left, right := 0, x
	ans := -1

	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func mySqrt3(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(xi-x0) < 1e-7 {
			break
		}
		x0 = xi
	}

	return int(x0)
}
