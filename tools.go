package main

import (
	_ "github.com/cheekybits/genny/generic"
	"sort"
)

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func power(x, y int) int {
	var ans int = 1
	for i := 1; i <= y; i++ {
		ans = ans * x
	}

	return ans
}

func getMaxArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func getMinArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}

func reverseString(s string) string {
	strArr := []byte(s)
	length := len(strArr)

	for i := 0; i < length/2; i++ {
		strArr[i], strArr[length-1-i] = strArr[length-1-i], strArr[i]
	}

	return string(strArr)
}

func equalOfTwoIntArray(array1, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}
