package abc

import (
	"math/rand"
)

func GenerateRandIntArr(n int) []int {
	scope := 100
	numbers := []int{}
	// var numbers []int

	for i := 0; i < n; i++ {
		// numbers[i] = rand.Intn(scope)
		numbers = append(numbers, rand.Intn(scope))
	}

	return numbers
}
