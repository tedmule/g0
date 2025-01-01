package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(err)
				return
			}
			message := string(buffer)
			fmt.Printf("Received: %s", message)
		}
	}()

	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}
