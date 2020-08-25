// Package lengthconv performs Feet and Meter computations.
package lengthconv

import (
	"testing"
)

func TestFToM(t *testing.T) {
	type args struct {
		f Feet
	}
	tests := []struct {
		name string
		args args
		want Meter
	}{
		{"length of marathon", args{Feet(42195 / 0.3048)}, MarathonM},
		{"height of Everest", args{Feet(8848 / 0.3048)}, EverestM},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToM(tt.args.f); got != tt.want {
				t.Errorf("FToM(Feet(%#v)) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}

func TestMToF(t *testing.T) {
	type args struct {
		f Meter
	}
	tests := []struct {
		name string
		args args
		want Feet
	}{
		{"length of marathon", args{MarathonM}, Feet(42195 / 0.3048)},
		{"height of Everest", args{EverestM}, Feet(8848 / 0.3048)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MToF(tt.args.f); got != tt.want {
				t.Errorf("MToF(Meter(%#v)) = %v, want %v", tt.args.f, got, tt.want)
			}
		})
	}
}

func TestFeet_String(t *testing.T) {
	got := Feet(10).String()
	want := "10ft"
	if got != want {
		t.Errorf("Feet(10).String() = %v, want %v", got, want)
	}
}

func TestMeter_String(t *testing.T) {
	got := Meter(10).String()
	want := "10m"
	if got != want {
		t.Errorf("Meter(10).String() = %v, want %v", got, want)
	}
}
