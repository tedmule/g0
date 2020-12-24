package sorting

func BubbleSort(numbers []int) []int {
	// Outer loop starts from length to 1,
	for i := len(numbers); i > 0; i-- {
		// Innger loop starts from 1,
		// so use j-1 for CHOOSING adjacent element.
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				tmp := numbers[j-1]
				numbers[j-1] = numbers[j]
				numbers[j] = tmp
			}
		}
	}
	return numbers
}
func BubbleSort2(numbers []int) []int {
	n := len(numbers)
	// Outer loop starts from 0 to length-1
	for i := 0; i < n; i++ {
		// Innter loop also starts from 0,
		// so we CHOOSE j+1 for the adjacent element,
		// and use n-i-1 to prevent array from OutOfIndex
		for j := 0; j < (n - i - 1); i++ {
			if numbers[j] > numbers[j+1] {
				tmp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = tmp
			}
		}

	}
	return numbers
}

func BubbleSortOptimized(numbers []int) {}
