package popcount

// pc[i] is the population count of i.
var _pc [256]byte

func init() {
	for i := range _pc {
		_pc[i] = _pc[i/2] + byte(i&1)
	}
}

// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(_pc[byte(x>>(i*8))])
	}
	return sum
}
