package abc

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 10)
	fmt.Printf("done")

	done <- true
}

// func main() {
// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done
// }
