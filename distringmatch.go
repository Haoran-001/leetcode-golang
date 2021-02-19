package main

func diStringMatch(S string) []int {
	length := len(S)
	ans := make([]int, length+1)
	low, high := 0, length
	sCharArr := []byte(S)

	for i := 0; i < length; i++ {
		if sCharArr[i] == 'I' {
			ans[i] = low
			low++
		} else if sCharArr[i] == 'D' {
			ans[i] = high
			high--
		}
	}

	ans[length] = low

	return ans
}
