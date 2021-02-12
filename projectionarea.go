package main

func projectionArea(grid [][]int) int {
	var count int = 0
	// 计算底部阴影面积, 数组中所有不为零的元素
	for _, item := range grid {
		for i := 0; i < len(item); i++ {
			if item[i] > 0 {
				count++
			}
		}
	}
	// 计算正面阴影面积, 所有行每行最大的元素值累加
	for _, item := range grid {
		maxValue := 0
		for i := 0; i < len(item); i++ {
			maxValue = max(maxValue, item[i])
		}
		count += maxValue
	}
	// 计算侧面阴影面积, 所有列每列最大的元素值累加
	for i := 0; i < len(grid[0]); i++ {
		maxValue := 0
		for j := 0; j < len(grid); j++ {
			maxValue = max(maxValue, grid[j][i])
		}
		count += maxValue
	}

	return count
}