package main

func maximum(a int, b int) int {
	diff := int64(a) - int64(b)
	// a和b的差值为大于等于0, 则sign返回0, 否则返回-1
	sign := diff >> 63

	return a*(int(sign)+1) - b*int(sign)
}
