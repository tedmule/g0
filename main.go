package main

import (
	nw "github.com/g0gogo/netswatch"
)

func main() {
	nw.Hello()
	// wg := sync.WaitGroup{}
	// ctx := context.Background()

	// nw.ListContainers(ctx)

	// wg.Add(1)
	// go func() {
	// 	nw.WatchNetwork(ctx, &wg)
	// }()

	// wg.Add(1)
	// go func() {
	// 	nw.WatchCtrEvents(ctx, &wg)
	// }()

	// wg.Wait()
}
