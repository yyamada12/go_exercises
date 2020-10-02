package main

import (
	"math/big"
	"testing"
)

func Test_const(t *testing.T) {
	if KB != 1000 {
		t.Errorf("KB got %26.f, want 1000", KB)
	}
	if MB != 1000000 {
		t.Errorf("MB got %26.f, want 1000000", MB)
	}
	if GB != 1000000000 {
		t.Errorf("GB got %26.f, want 1000000000", GB)
	}
	if TB != 1000000000000 {
		t.Errorf("TB got %26.f, want 1000000000000", TB)
	}
	if PB != 1000000000000000 {
		t.Errorf("PB got %26.f, want 1000000000000000", PB)
	}
	if EB != 1000000000000000000 {
		t.Errorf("EB got %26.f, want 1000000000000000000", EB)
	}
	if ZB != 1000000000000000000000 {
		t.Errorf("ZB got %26.f, want 1000000000000000000000", ZB)
	}
	if YB != 1000000000000000000000000 {
		t.Errorf("YB got %d, want 1000000000000000000000000", new(big.Int).Mul(big.NewInt(YB/TB), big.NewInt(TB)))
	}
}
