package sorting

func BubbleSort(numbers []int) []int {

	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				tmp := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = tmp
			}
		}
	}
	return numbers
}

func BubbleSortOptimized(numbers []int) {}
