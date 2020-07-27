package main

import (
	"fmt"
	"time"

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
	var dnsRegistry nw.DNSRegistry

	dnsRegistry.Endpoint = "http://172.16.66.10:8500"
	dnsRegistry.Token = "cd3895fe-04d6-609e-4e3d-138e5a0bbf3b"

	dnsRegistry.RegisterSvc()
	time.Sleep(2 * time.Second)
	dnsRegistry.ListService()

	meta := nw.GenerateNodeMeta()
	fmt.Printf("%+v", meta)
}
