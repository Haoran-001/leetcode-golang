package main

import "fmt"

func main() {
	//const UINT_MIN uint = 0
	//const UINT_MAX uint = ^uint(0)
	//const INT_MAX int = int(^uint(0)>>1)
	//const INT_MIN int = ^INT_MAX
	//fmt.Println(UINT_MAX)
	//fmt.Println(UINT_MIN)
	//fmt.Println(INT_MAX)
	//fmt.Println(INT_MIN)

	var grid [][]int = [][]int {{1, 2}, {3, 4}}
	fmt.Println(projectionArea(grid))

	var	points[][]int = [][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}
	fmt.Println(largestTriangleArea(points))
	var grid2 [][]int = [][]int{{1, 2}, {3, 4}}
	fmt.Println(surfaceArea2(grid2))
}
