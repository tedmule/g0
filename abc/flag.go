package abc

import (
	"flag"
	"fmt"
)

//var ip = flag.Int("flagname", 1024, "help message for t66y")
var flagvar int

func DemoFlags() {
	flag.IntVar(&flagvar, "flagname", 1024, "help me")
	flag.Parse()
	fmt.Println("ip:", flagvar)
}
