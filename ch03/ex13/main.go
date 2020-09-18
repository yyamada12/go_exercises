package main

import "fmt"

func main() {
	const (
		KB = 1e3
		MB = 1e6
		GB = 1e9
		TB = 1e12
		PB = 1e15
		EB = 1e18
		ZB = 1e21
		YB = 1e24
	)

	fmt.Printf("KB = %26.f\n", KB)
	fmt.Printf("MB = %26.f\n", MB)
	fmt.Printf("GB = %26.f\n", GB)
	fmt.Printf("TB = %26.f\n", TB)
	fmt.Printf("PB = %26.f\n", PB)
	fmt.Printf("EB = %26.f\n", EB)
	fmt.Printf("ZB = %26.f\n", ZB)
	fmt.Printf("YB = %26.f\n", YB)
}
