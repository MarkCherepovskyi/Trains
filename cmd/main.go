package main

import "DL/trains/pkg"

func main() {
	info := pkg.ModelInit()
	pkg.InitTSP(info)
	pkg.Do()
}
