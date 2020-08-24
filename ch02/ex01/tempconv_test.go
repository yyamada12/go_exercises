// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{"freezing temp", args{FreezingC}, Fahrenheit(32)},
		{"boiling temp", args{BoilingC}, Fahrenheit(212)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToF(tt.args.c); got != tt.want {
				t.Errorf("CToF(Celsius(%#v)) = %v, want %v", tt.args.c, got, tt.want)
			}
		})
	}
}

func TestCToK(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{"freezing temp", args{FreezingC}, Kelvin(273.15)},
		{"boiling temp", args{BoilingC}, Kelvin(373.15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.args.c); got != tt.want {
				t.Errorf("CToK(Celsius(%#v)) = %v, want %v", tt.args.c, got, tt.want)
			}
		})
	}
}

func TestFToK(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{"freezing temp", args{Fahrenheit(32)}, Kelvin(273.15)},
		{"boiling temp", args{Fahrenheit(212)}, Kelvin(373.15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToK(tt.args.f); got != tt.want {
				t.Errorf("FToK(Fahrenheit(%#v)) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{"freezing temp", args{Fahrenheit(32)}, FreezingC},
		{"boiling temp", args{Fahrenheit(212)}, BoilingC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToC(tt.args.f); got != tt.want {
				t.Errorf("FToC(Fahrenheit(%#v)) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}

func TestKToC(t *testing.T) {
	type args struct {
		f Kelvin
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{"freezing temp", args{Kelvin(273.15)}, FreezingC},
		{"boiling temp", args{Kelvin(373.15)}, BoilingC},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToC(tt.args.f); got != tt.want {
				t.Errorf("KToC(Kelvin(%#v)) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}

func TestKToF(t *testing.T) {
	type args struct {
		c Kelvin
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{"freezing temp", args{Kelvin(273.15)}, Fahrenheit(32)},
		{"boiling temp", args{Kelvin(373.15)}, Fahrenheit(212)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToF(tt.args.c); got != tt.want {
				t.Errorf("KToF(Kelvin(%#v)) = %v, want %v", tt.args.c, got, tt.want)
			}
		})
	}
}

func TestCelsius_String(t *testing.T) {
	got := Celsius(10).String()
	want := "10°C"
	if got != want {
		t.Errorf("Celsius(10).String() = %v, want %v", got, want)
	}
}

func TestFahrenheit_String(t *testing.T) {
	got := Fahrenheit(10).String()
	want := "10°F"
	if got != want {
		t.Errorf("Fahrenheit(10).String() = %v, want %v", got, want)
	}
}

func TestKelvin_String(t *testing.T) {
	got := Kelvin(10).String()
	want := "10K"
	if got != want {
		t.Errorf("Kelvin(10).String() = %v, want %v", got, want)
	}
}
