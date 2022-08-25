package main

import (
	"DL/trains/pkg"
	"context"
	"sync"
)

func main() {
	info := pkg.ModelInit()

	pkg.InitTSP(info)

	pkg.InitTSP2(info)
	context.Background()
	wt := sync.WaitGroup{}
	wt.Add(2)
	go pkg.DoTSPbyTime(&wt)
	go pkg.Do(&wt)
	wt.Wait()

}
