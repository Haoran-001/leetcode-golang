package main

import "sort"

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

func getMaxArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func getMinArrayElements(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}
