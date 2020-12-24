package abc

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
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

func CalFileSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(s string, info os.FileInfo, err error) error {
		fmt.Println(s)
		fmt.Println(info, info.Name())
		fmt.Println(err)
		fmt.Println("-------------")
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(info.Name())
			size += info.Size()
		}
		return err
	})
	return size, err
}
