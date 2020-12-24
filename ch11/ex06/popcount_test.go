package popcount

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var output int

func BenchPopCountOriginal(b *testing.B, input uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(input)
	}
	output = s
}

func BenchPopCountBitShift(b *testing.B, input uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountBitShift(input)
	}
	output = s
}

func BenchPopCountBitClear(b *testing.B, input uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountBitClear(input)
	}
	output = s
}

func BenchmarkPopCount(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for bitLength := 8; bitLength < 64; bitLength += 8 {
		input := rng.Int63n(1 << bitLength)
		b.Run(fmt.Sprintf("Original:bitLength=%d", bitLength), func(b *testing.B) {
			BenchPopCountOriginal(b, uint64(input))
		})
		b.Run(fmt.Sprintf("BitShift:bit=%d", bitLength), func(b *testing.B) {
			BenchPopCountBitShift(b, uint64(input))
		})
		b.Run(fmt.Sprintf("BitClear:bit=%d", bitLength), func(b *testing.B) {
			BenchPopCountBitClear(b, uint64(input))
		})
	}
}
