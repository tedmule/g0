package sorting

import (
	"math/rand"
)

func QuickSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	left, right := 0, len(numbers)-1

	// Pick a pivot
	pivot := rand.Int() % len(numbers)
	// Exchange the pivot with last element
	// or comment out the line below, make the last element be a pivot always
	numbers[pivot], numbers[right] = numbers[right], numbers[pivot]

	for i := range numbers {
		if numbers[i] < numbers[right] {
			numbers[left], numbers[i] = numbers[i], numbers[left]
			left++
		}
	}
	// Exchange element at index left with element at index right
	numbers[left], numbers[right] = numbers[right], numbers[left]

	// Then partition the array with pivot(the element at index right)
	QuickSort(numbers[:left])
	QuickSort(numbers[left+1:])

	return numbers
}

func QuickSort2(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	left, right := 0, len(numbers)-1
	pivot := numbers[right]
	partitionIdx := left

	for i := range numbers {
		if numbers[i] < pivot {
			numbers[i], numbers[partitionIdx] = numbers[partitionIdx], numbers[i]
			partitionIdx++
		}
	}

	// Exchange pivot(the last element) with partitionIdx(left)
	numbers[partitionIdx], numbers[right] = numbers[right], numbers[partitionIdx]

	QuickSort2(numbers[:partitionIdx])
	QuickSort2(numbers[partitionIdx+1:])
	return numbers
}
