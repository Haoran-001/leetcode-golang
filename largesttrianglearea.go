package main

import (
	"math"
)


func largestTriangleArea(points [][]int) float64 {
	var maxArea float64 = 0.0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				curArea := area(points[i], points[j], points[k])
				maxArea = math.Max(maxArea, curArea);
			}
		}
	}
	return maxArea
}

func largestTriangleArea2(points [][]int) float64 {
	var maxArea float64 = 0

	for i:=0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				curArea := area2(points[i], points[j], points[k])
				maxArea = math.Max(maxArea, curArea)
			}
		}
	}
	return maxArea
}

func largestTriangleArea3(points [][]int) float64 {
	var maxArea float64 = 0

	for i:=0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				curArea := area3(points[i], points[j], points[k])
				maxArea = math.Max(maxArea, curArea)
			}
		}
	}
	return maxArea
}

// 鞋带定理
func area(point1 []int, point2 []int, point3 []int) float64 {
	var area float64 = 0.0
	var a int = point1[0] * point2[1] + point2[0] * point3[1] + point3[0] * point1[1];
	var b int = point1[1] * point2[0] + point2[1] * point3[0] + point3[1] * point1[0];
	area = 0.5 * math.Abs(float64(a) - float64(b))

	return area
}

// 海伦定理
func area2(point1 []int, point2 []int, point3 []int) float64 {
	x1 := float64(point1[0]); y1 := float64(point1[1])
	x2 := float64(point2[0]); y2 := float64(point2[1])
	x3 := float64(point3[0]); y3 := float64(point3[1])

	a := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))
	b := math.Sqrt((x3 - x2) * (x3 - x2) + (y3 - y2) * (y3 - y2))
	c := math.Sqrt((x1 - x3) * (x1 - x3) + (y1 - y3) * (y1 - y3))

	C := (a + b + c) / 2

	S := math.Sqrt(C * (C - a) * (C - b) * (C - c))

	return S
}

// 1/2*a*b*sin(c)
func area3(point1 []int, point2 []int, point3 []int) float64 {
	x1 := float64(point1[0]); y1 := float64(point1[1])
	x2 := float64(point2[0]); y2 := float64(point2[1])
	x3 := float64(point3[0]); y3 := float64(point3[1])

	a := math.Sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))
	b := math.Sqrt((x3 - x2) * (x3 - x2) + (y3 - y2) * (y3 - y2))
	c := math.Sqrt((x1 - x3) * (x1 - x3) + (y1 - y3) * (y1 - y3))

	cosc := (a * a + b * b - c * c) / 2 * a * b
	sinc := math.Sqrt(1 - cosc * cosc)

	S := 0.5 * a * b * sinc

	return S
}
