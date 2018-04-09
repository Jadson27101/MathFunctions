package main

import "fmt"

func mulGF256(a uint64, b uint64) uint64 {
	var result uint64 = 0
	var counter uint64
	var hi_bit uint64
	for counter = 0; counter < 8; counter++ {
		if (b & 1) != 0 {
			result ^= a
		}
		hi_bit = a & 0x80
		a <<= 1
		if hi_bit != 0 {
			a ^= 0x11b
		}
		b >>= 1
	}
	return result
}
func main()  {
	fmt.Println(mulGF256(6,24))
}
