package popcount

// PopCountBitClear returns the population count (number of set bits) of x.
func PopCountBitClear(x uint64) int {
	sum := 0
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}
