package main

import (
	"context"
	"sync"

	nw "github.com/g0gogo/netswatch"
	"github.com/g0gogo/packs"
)

func main() {
	nw.Hello()
	wg := sync.WaitGroup{}
	ctx := context.Background()

	nw.ListContainers(ctx)
	packs.DemoJson()

	// wg.Add(1)
	// go func() {
	// 	nw.WatchNetwork(ctx, &wg)
	// }()

	// wg.Add(1)
	// go func() {
	// 	nw.WatchCtrEvents(ctx, &wg)
	// }()

	wg.Wait()
}
