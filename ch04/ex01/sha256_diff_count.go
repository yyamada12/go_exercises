package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

var pc [256]byte

const size = 32

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, `USAGE: go run sha256_diff_count.go digest1 digest2`)
	}
	arg1, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decode err: %s is not hex\n", os.Args[1])
	}
	if len(arg1) != size {
		fmt.Fprintf(os.Stderr, "Value err: %s is not digest of sha256\n", os.Args[1])
	}
	arg2, err := hex.DecodeString(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decode err: %s is not hex\n", os.Args[2])
	}
	if len(arg2) != size {
		fmt.Fprintf(os.Stderr, "Value err: %s is not digest of sha256\n", os.Args[2])
	}
	var d1, d2 [size]byte
	copy(d1[:], arg1[:])
	copy(d2[:], arg2[:])
	fmt.Println(diffCount(d1, d2))
}

func diffCount(d1, d2 [size]byte) byte {
	c := byte(0)
	for i := 0; i < size; i++ {
		c += pc[d1[i]^d2[i]]
	}
	return c
}
