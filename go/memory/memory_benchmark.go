package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	N := 10_000_000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i * 2
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("Memory Benchmark: Sum: %d | Time: %v\n", sum, time.Since(start))
}
