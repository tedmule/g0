package sorting

func SelectionSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		// SELECT i as the fake minium number
		min := i
		//Inner loop to scan the real mininum number
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		// When the inner loop is done,
		// swap the fake minium number with real minium number
		numbers[i], numbers[min] = numbers[min], numbers[i]
	}
	return numbers
}
