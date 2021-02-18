package main

const MOD int = 1e9 + 7

func numPrimeArrangements(n int) int {
	countOfPrime, countOfNoPrime := 0, 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			countOfPrime++
		}
	}

	countOfNoPrime = n - countOfPrime

	ans := (factorial(countOfPrime) * factorial(countOfNoPrime)) % MOD

	return ans
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// mod 10^9 + 7
func factorial(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans = (ans * i) % MOD
	}

	return ans
}
