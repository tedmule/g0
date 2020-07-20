package main

import (
    "fmt"
    "flag"
)

//var ip = flag.Int("flagname", 1024, "help message for t66y")
var flagvar int

func main() {
    flag.IntVar(&flagvar, "flagname", 1024, "help me")
    flag.Parse()
    fmt.Println("ip:", flagvar)
}
