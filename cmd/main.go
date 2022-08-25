package main

import (
	"DL/trains/pkg"
)

func main() {
	info := pkg.ModelInit()

	pkg.InitTSP(info)

	//pkg.InitTSP2(info)

	//go pkg.DoTSPbyTime(&wt)
	pkg.Do()

}
