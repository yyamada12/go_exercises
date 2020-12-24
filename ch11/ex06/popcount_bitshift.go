package popcount

// PopCountBitShift returns the population count (number of set bits) of x.
func PopCountBitShift(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		sum += int((x >> i) & 1)
	}
	return sum
}
