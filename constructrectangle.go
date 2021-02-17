package main

import "math"

func constructRectangle(area int) []int {
	areaFloat64 := float64(area)
	length := math.Sqrt(areaFloat64)

	if int(length)*int(length) == area {
		return []int{int(length), int(length)}
	}

	for i := int(length); i >= 1; i-- {
		if i*(area/i) == area {
			return []int{area / i, i}
		}
	}

	return []int{0, 0}
}
