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
	sorted := sorting.BubbleSort(arr)
	fmt.Printf("Sorted: %v\n", sorted)

	// arr2 := abc.GenerateRandIntArr(10)
	// fmt.Printf("Origin: %+v\n", arr2)
	// defer abc.Elapsed()()
	// sorted2 := sorting.BubbleSort2(arr2)
	// fmt.Printf("Sorted: %v\n", sorted2)
}
