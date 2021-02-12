package main


func surfaceArea(grid [][]int) int {
	var area int = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			level := grid[i][j]
			if level > 0 {
				area += level * 4 + 2
				if i > 0 {
					area -= min(level, grid[i - 1][j]) * 2
				}
				if j > 0 {
					area -= min(level, grid[i][j - 1]) * 2
				}
			}
		}
	}

	return area
}

func surfaceArea2(grid [][]int) int {
	var area int = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 不考虑柱体间侧面积重叠部分, 计算每个柱体的表面积
			if grid[i][j] > 0 {
				area += (grid[i][j] * 4 + 2)
			}

			// 计算i和i-1侧面积贴合的部分, 需要减掉
			if i >= 1 {
				area -= (min(grid[i][j], grid[i - 1][j]) * 2)
			}

			// 计算j和j-1侧面积贴合的部分, 需要减掉
			if j >= 1{
				area -= (min(grid[i][j], grid[i][j - 1]) * 2)
			}
		}
	}
	return area
}