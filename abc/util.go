package abc

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandIntArr(n int) []int {
	rand.Seed(time.Now().Unix())
	scope := 100
	numbers := []int{}

	for i := 0; i < n; i++ {
		numbers = append(numbers, rand.Intn(scope))
	}

	return numbers
}

func Elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time: %v\n", time.Since(start))
	}
}
