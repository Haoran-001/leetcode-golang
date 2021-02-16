package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 如果两个矩形至少有一个为0, 则显然是不重叠
	if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
		return false
	}

	// 判断不重叠的情况, 即在两个矩形的上下左右任意一个方位存在一条竖线可以将其两个矩形分开
	// 对不重叠的情况取否命题即可
	isNotOverlapping := (rec1[3] <= rec2[1]) || (rec1[1] >= rec2[3]) || (rec1[2] <= rec2[0]) || (rec1[0] >= rec2[2])

	return !isNotOverlapping
}

func isRectangleOverlap2(rec1 []int, rec2 []int) bool {
	// 将两个矩形的x坐标和y坐标分别投影到x轴上和y轴上, 在x轴和y轴上分别形成两条线段, 共四条线段
	// 若x轴上的两条线段有交集且y轴上的两条线段也有交集, 则表示两个矩形重叠
	// 矩形1在x轴上的投影线段为(rec1[0], rec1[2]), 在y轴上的投影为(rec1[1], rec[3])
	// 矩形2在x轴上的投影线段为(rec2[0], rec2[2]), 在y轴上的投影为(rec2[1], rec2[3])
	overlapping := max(rec1[0], rec2[0]) < min(rec1[2], rec2[2]) && max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])

	return overlapping
}
