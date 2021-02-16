package main

func isBoomerang(points [][]int) bool {
	// 若存在至少两个点相同, 则返回false
	if equal(points[0], points[1]) || equal(points[0], points[2]) || equal(points[1], points[2]) {
		return false
	}

	// 若三个点的横坐标都相等, 则返回false
	if points[0][0] == points[1][0] && points[1][0] == points[2][0] {
		return false
	}

	k1 := slope(points[0], points[1])
	k2 := slope(points[0], points[2])

	return k1 != k2
}

func slope(point1, point2 []int) float64 {
	x1, y1 := float64(point1[0]), float64(point1[1])
	x2, y2 := float64(point2[0]), float64(point2[1])

	k := (y2 - y1) / (x2 - x1)

	return k
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
