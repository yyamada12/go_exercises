package bigcomplex

import (
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkZ3Rat(b *testing.B) {
	const denom = 100000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		x := big.NewRat(rand.Int63()%(denom*2), denom)
		y := big.NewRat(rand.Int63()%(denom*2), denom)
		if rand.Int()%2 == 0 {
			x = new(big.Rat).Neg(x)
		}
		if rand.Int()%2 == 0 {
			y = new(big.Rat).Neg(y)
		}
		z := NewCmplRat(x, y)
		z = z.Times(z).Times(z)
	}
}

func BenchmarkZ4Rat(b *testing.B) {
	const denom = 100000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		x := big.NewRat(rand.Int63()%(denom*2), denom)
		y := big.NewRat(rand.Int63()%(denom*2), denom)
		if rand.Int()%2 == 0 {
			x = new(big.Rat).Neg(x)
		}
		if rand.Int()%2 == 0 {
			y = new(big.Rat).Neg(y)
		}
		z := NewCmplRat(x, y)
		z = z.Times(z).Times(z).Times(z)
	}
}

func Benchmark1dZRat(b *testing.B) {
	const denom = 100000
	rand.Seed(time.Now().UnixNano())
	one := NewCmplRat(big.NewRat(1, 1), big.NewRat(0, 1))
	for i := 0; i < b.N; i++ {
		x := big.NewRat(rand.Int63()%(denom*2), denom)
		y := big.NewRat(rand.Int63()%(denom*2), denom)
		if rand.Int()%2 == 0 {
			x = new(big.Rat).Neg(x)
		}
		if rand.Int()%2 == 0 {
			y = new(big.Rat).Neg(y)
		}
		z := NewCmplRat(x, y)
		z = one.Divides(z)
	}
}

func BenchmarkCmpZRat(b *testing.B) {
	const denom = 100000
	rand.Seed(time.Now().UnixNano())
	one := NewCmplRat(big.NewRat(1, 1), big.NewRat(0, 1))
	for i := 0; i < b.N; i++ {
		x := big.NewRat(rand.Int63()%(denom*2), denom)
		y := big.NewRat(rand.Int63()%(denom*2), denom)
		if rand.Int()%2 == 0 {
			x = new(big.Rat).Neg(x)
		}
		if rand.Int()%2 == 0 {
			y = new(big.Rat).Neg(y)
		}
		z := NewCmplRat(x, y)
		if z.SquaredAbs().Cmp(big.NewRat(1, 1e12)) < 0 {
			z = one
		}
	}
}
