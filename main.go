package main

import (
	"bytes"
	"fmt"
	"sort"
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
	fmt.Println("hello")
	fmt.Println(getNoZeroIntegers(12))
	fmt.Println(subtractProductAndSum2(13))
	fmt.Println(isPowerOfTwo2(1024))
	var num int32 = 1 << 30
	fmt.Println(num)
	fmt.Println(printNumbers(3))
	fmt.Println(addDigits2(38))
	fmt.Println(totalMoney2(20))
	points3 := [][]int{{1, 2}, {1, 3}, {1, 4}}
	fmt.Println(isBoomerang(points3))
	arr4 := []int{3, 4, 2, 1}
	sort.Ints(arr4)
	fmt.Println(arr4)
	rec1, rec2 := []int{0, 0, 1, 1}, []int{1, 0, 2, 1}
	fmt.Println(isRectangleOverlap(rec1, rec2))
	sort.Sort(sort.Reverse(sort.IntSlice(arr4)))
	fmt.Println(arr4)

	stu := []Student{
		{Name: "haoran", Age: 24},
		{Name: "xiaorong", Age: 27},
		{Name: "mama", Age: 50},
		{"papa", 49},
	}

	sort.Sort(StudentSlice(stu))
	fmt.Println(stu)
	sort.Sort(sort.Reverse(StudentSlice(stu)))
	fmt.Println(stu)

	fmt.Println(isPowerOfThree3(27))
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(maxCount2(3, 3, [][]int{{2, 2}, {3, 3}}))

	fmt.Println(-100 >> 31)
	fmt.Println(maximum69Number(9669))

	fmt.Println(numPrimeArrangements(100))

	fmt.Println(binaryGap(22))
}
