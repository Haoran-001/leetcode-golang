package main

import "bytes"

func reverseStr(s string) string {
	sArr := []rune(s)
	length := len(sArr)
	for i := 0; i < length/2; i++ {
		sArr[i], sArr[length-1-i] = sArr[length-1-i], sArr[i]
	}

	return string(sArr)
}

func convertToTitle(n int) string {
	var buffer bytes.Buffer
	for {
		n = n - 1
		i := byte(n % 26)
		buffer.WriteByte('A' + i)
		n = n / 26

		if n == 0 {
			break
		}
	}

	return reverseStr(buffer.String())
}
