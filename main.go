package main

import (
	"bytes"
	"fmt"
)

func main() {
	//const UINT_MIN uint = 0
	//const UINT_MAX uint = ^uint(0)
	//const INT_MAX int = int(^uint(0)>>1)
	//const INT_MIN int = ^INT_MAX
	//fmt.Println(UINT_MAX)
	//fmt.Println(UINT_MIN)
	//fmt.Println(INT_MAX)
	//fmt.Println(INT_MIN)

	var grid [][]int = [][]int{{1, 2}, {3, 4}}
	fmt.Println(projectionArea(grid))

	var points [][]int = [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	fmt.Println(largestTriangleArea(points))
	var grid2 [][]int = [][]int{{1, 2}, {3, 4}}
	fmt.Println(surfaceArea2(grid2))
	fmt.Println(titleToNumber("ZYY"))
	fmt.Println(18 % 4)
	var buffer bytes.Buffer
	buffer.WriteByte('A')
	buffer.WriteByte('B')
	fmt.Println(convertToTitle(26))
	fmt.Println(mySqrt3(8))
	fmt.Println(checkPerfectNumber2(28))
	fmt.Println(arrangeCoins(1))
	fmt.Println(checkPerfectNumber3(8128))
}
