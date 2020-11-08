package main

import (
    "fmt"
)


func main() {
    //months := [...]string{1: "Jan", 2: "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
    //a := months[:12]
    //b := months[:]
    //fmt.Println(a)
    //fmt.Println(b)
    a := [...]string{1: "apple", 2: "banana", 3: "orange"}

    for i, val := range a {
        fmt.Println(i, val)
    }
    fmt.Println(len(a))
}
