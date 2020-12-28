package intset

import (
	"math/rand"
	"testing"
	"time"
)

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

func Bench_Add(b *testing.B, size, max int) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	var inputs []int
	for i := 0; i < size; i++ {
		inputs = append(inputs, rng.Intn(max))
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

func Bench_UnionWith(b *testing.B, size, max int) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	var x, y IntSet
	x2 := NewMapIntSet()
	y2 := NewMapIntSet()
	for i := 0; i < size; i++ {
		r1 := rng.Intn(max)
		r2 := rng.Intn(max)
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

func words_Has(b *testing.B, x *IntSet, e int) {
	for i := 0; i < b.N; i++ {
		x.Has(e)
	}
}

func map_Has(b *testing.B, x *MapIntSet, e int) {
	for i := 0; i < b.N; i++ {
		x.Has(e)
	}
}

func Bench_Has(b *testing.B, size, max int) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	var x IntSet
	x2 := NewMapIntSet()
	for i := 0; i < size; i++ {
		r := rng.Intn(max)
		x.Add(r)
		x2.Add(r)
	}
	e := rng.Intn(max)

	b.Run("words", func(b *testing.B) {
		words_Has(b, &x, e)
	})
	b.Run("map", func(b *testing.B) {
		map_Has(b, x2, e)
	})
}

func Benchmark_wordsAndMap(b *testing.B) {
	for _, size := range []int{10, 1000, 100000} {
		for _, max := range []int{100, 10000, 1000000} {
			b.Logf("size of set: %d", size)
			b.Logf("max element size: %d", max)
			b.Run("Add", func(b *testing.B) {
				Bench_Add(b, size, max)
			})
			b.Run("UnionWith", func(b *testing.B) {
				Bench_UnionWith(b, size, max)
			})
			b.Run("Has", func(b *testing.B) {
				Bench_Has(b, size, max)
			})
		}
	}
}
