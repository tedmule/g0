package main

import (
	"fmt"

	"github.com/cakturk/go-netstat/netstat"
)

var stat map[string]int

func main() {
	// TCP sockets
	socks, err := netstat.TCPSocks(netstat.NoopFilter)
	if err != nil {
		panic(err.Error())
	}
	for _, tcp := range socks {
		// fmt.Printf("%v\n", tcp)
		fmt.Println(tcp.RemoteAddr.IP)
		fmt.Println(tcp.RemoteAddr.Port)
		fmt.Println(tcp.State)
		fmt.Println(tcp.State)
	}

}

//val, ok := myMap["foo"]
//// If the key exists
//if ok {
//    // Do something
//}
