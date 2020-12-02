package main

import "math"

func isPrime(a int) int {
	if a == 1 {
		return 1
	}
	for i:=2; i<int(math.Sqrt(float64(a))); i++ {
		if a % i == 0 {
			return 0
		}
	}
	return 1
}

func countPrimes(n int) int {
	count := 0
	for i:=2; i<n; i++ {
		count += isPrime(i)
	}
	return count
}

func countPrimes2(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
