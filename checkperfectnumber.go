package main

func checkPerfectNumber(num int) bool {
	ans := 0

	for i := 1; i < num; i++ {
		if num%i == 0 {
			ans += i
		}
	}

	return num == ans
}

func checkPerfectNumber2(num int) bool {
	ans := 0

	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			ans += i
			if i*i != num {
				ans += (num / i)
			}
		}
	}

	return num*2 == ans
}

func checkPerfectNumber3(num int) bool {
	primeArray := []int{2, 3, 5, 7, 13, 17}
	// 由于目前的完美数都是偶完美数, 且由欧几里得-欧拉定理可得每个偶完全数可以写做: 2^(p-1)*(2^p-1)形式, 且这里的p需满足p本身是素数, 2^p-1也为素数
	// 显然, 11虽然为素数, 但2^11-1并不是素数, 所以p=11不满足要求
	// 题意的num取值范围在10^8以内, 因此只需要枚举到2, 3, 5, 7, 13, 17即可
	length := len(primeArray)
	ansArray := make([]int, length)

	for i := 0; i < length; i++ {
		ansArray[i] = (1 << (primeArray[i] - 1)) * (1<<primeArray[i] - 1)
	}

	for i := 0; i < length; i++ {
		if num == ansArray[i] {
			return true
		}
	}

	return false
}
