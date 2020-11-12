package main

import (
	"errors"
	"fmt"

	"github.com/g0gogo/packs"
)

func main() {
	filename := "/tmp/test"
	size, err := packs.CalFileSize(filename)

	if err != nil {
		errors.New("calculate error")
	} else {
		fmt.Println(size)
	}
}
