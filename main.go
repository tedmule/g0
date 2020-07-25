package main

import (
	"context"
	"sync"

	"github.com/g0gogo/netswatch"
	"github.com/g0gogo/packs"
)

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()

	netswatch.ListContainers(ctx)
	packs.DemoJson()

	wg.Add(1)
	go func() {
		netswatch.WatchNetwork(ctx, &wg)
	}()

	wg.Add(1)
	go func() {
		netswatch.WatchCtrEvents(ctx, &wg)
	}()
	wg.Wait()
}
