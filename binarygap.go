package main

import "strconv"

// 0010011
func binaryGap(n int) int {
	numToStr2Base := strconv.FormatInt(int64(n), 2)
	numCharArr := []byte(numToStr2Base)
	maxDistance, curDistance := 0, 0
	startIndex := -1

	// 记录第一个1的起始下标
	for i := 0; i < len(numCharArr); i++ {
		if numCharArr[i] == '1' {
			startIndex = i
			curDistance++
			break
		}
	}

	for i := startIndex + 1; i < len(numCharArr); i++ {
		if numCharArr[i] != '1' {
			curDistance++
		} else if numCharArr[i] == '1' {
			maxDistance = max(maxDistance, curDistance)
			curDistance = 1
		}
	}

	return maxDistance
}

func binaryGap2(n int) int {
	arr := [32]int{}
	index := 0
	maxDistance := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			arr[index] = i
			index++
		}
	}

	for i := 0; i < index; i++ {
		maxDistance = max(maxDistance, arr[i+1]-arr[i])
	}

	return maxDistance
}

func binaryGap3(n int) int {
	last := -1
	maxDistance := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			// 不是转换为二进制后右边第一个1的情况
			if last != -1 {
				maxDistance = max(maxDistance, i-last)
			}
			last = i
		}
	}

	return maxDistance
}
