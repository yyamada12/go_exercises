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
		var x IntSet
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
		words_Add(b, inputs)
	})

}

func Test_wordsAndMap(t *testing.T) {
	var x, y IntSet
	x2 := NewMapIntSet()
	y2 := NewMapIntSet()
	x.Add(1)
	x2.Add(1)
	x.Add(144)
	x2.Add(144)
	x.Add(9)
	x2.Add(9)
	if x.String() != x2.String() {
		t.Errorf("x got %s, want %s", x.String(), x2.String())
	}

	y.Add(9)
	y2.Add(9)
	y.Add(42)
	y2.Add(42)

	x.UnionWith(&y)
	x2.UnionWith(y2)
	if x.String() != x2.String() {
		t.Errorf("x union with y got %s, want %s", x.String(), x2.String())
	}

	if x.Has(9) != x2.Has(9) {
		t.Errorf("x.Has(9) got %t, want %t", x.Has(9), x2.Has(9))
	}
	if x.Has(123) != x2.Has(123) {
		t.Errorf("x.Has(123) got %t, want %t", x.Has(123), x2.Has(123))
	}
}
