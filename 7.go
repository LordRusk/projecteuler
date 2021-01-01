// find the 10001 prime
package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	
	return true
}

func main() {
	primes := []int{1, 2, 5, 7}
	for i := 8; len(primes) < 10001; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Printf("%v\n", primes[len(primes)-1])
}
