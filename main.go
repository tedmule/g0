package main

import (
	"fmt"

	"github.com/gog0/abc"
	"github.com/gog0/algo/sorting"
)

func main() {
	arr := abc.GenerateRandIntArr(10)

	fmt.Printf("Origin: %+v\n", arr)

	defer abc.Elapsed()()

	// sorted := sorting.BubbleSort(arr)
	// sorted := sorting.QuickSort(arr)
	sorted := sorting.QuickSort2(arr)

	fmt.Printf("Sorted: %v\n", sorted)
}
