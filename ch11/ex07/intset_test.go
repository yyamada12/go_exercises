package intset

import (
	"math/rand"
	"testing"
	"time"
)

var out IntSet

func words_Add(b *testing.B, inputs []int) {
	for i := 0; i < b.N; i++ {
		var x IntSet
		for _, input := range inputs {
			x.Add(input)
		}
	}
}

func map_Add(b *testing.B, inputs []int) {
	for i := 0; i < b.N; i++ {
		x := NewMapIntSet()
		for _, input := range inputs {
			x.Add(input)
		}
	}
}

func Benchmark_Add(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	var inputs []int
	for i := 0; i < 100000; i++ {
		inputs = append(inputs, rng.Intn(10000))
	}
	b.Run("words", func(b *testing.B) {
		words_Add(b, inputs)
	})
	b.Run("map", func(b *testing.B) {
		map_Add(b, inputs)
	})
}

func words_UnionWith(b *testing.B, x, y *IntSet) {
	for i := 0; i < b.N; i++ {
		x.UnionWith(y)
	}
}

func map_UnionWith(b *testing.B, x, y *MapIntSet) {
	for i := 0; i < b.N; i++ {
		x.UnionWith(y)
	}
}

func Benchmark_UnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	var x, y IntSet
	x2 := NewMapIntSet()
	y2 := NewMapIntSet()
	for i := 0; i < 100000; i++ {
		r1 := rng.Intn(10000)
		r2 := rng.Intn(10000)
		x.Add(r1)
		y.Add(r2)
		x2.Add(r1)
		y2.Add(r2)
	}
	b.Run("words", func(b *testing.B) {
		words_UnionWith(b, &x, &y)
	})
	b.Run("map", func(b *testing.B) {
		map_UnionWith(b, x2, y2)
	})
}
