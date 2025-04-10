// cpu_benchmark.go
package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	count := 0
	N := 1_000_000
	for i := 0; i < N; i++ {
		if isPrime(i) {
			count++
		}
	}
	fmt.Printf("Found %d primes in %dms\n", count, time.Since(start).Milliseconds())
}
