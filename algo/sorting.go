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
        // Inner loop also starts from 0,
        // so we CHOOSE j+1 for the adjacent element,
        // and use n-i-1 to prevent array from OutOfIndex
        for j := 0; j < (n - i - 1); j++ {
            if numbers[j] > numbers[j+1] {
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }   
        }   
        // fmt.Println(numbers)
    }   
    return numbers
}

func BubbleSortOptimized(numbers []int) []int {
    n := len(numbers)
    for i := 0; i < n; i++ {
        swapped := false
        for j := 0; j < (n - i - 1); j++ {
            if numbers[j] > numbers[j+1] {
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
                swapped = true
            }   
        }   
        // fmt.Println(numbers)
        if !swapped {
            break
        }   
    }   
    return numbers
}

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
