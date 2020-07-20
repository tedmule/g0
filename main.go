package main

import (
    "fmt"
)

func main() {
    i := 3
    pi := &i

    fmt.Printf("%v\n", i)
    fmt.Printf("%v\n", *pi)
}
