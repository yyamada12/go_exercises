// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"testing"

	"github.com/yyamada12/go_exercises/ch07/ex06/tempconv"
)

func Test_Celsius(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want float64
	}{
		{"default", "", 20.0},
		{"10C", "10C", 10.0},
		{"30°C", "30°C", 30.0},
		{"50F", "50F", 10.0},
		{"59°F", "59°F", 15.0},
		{"283.15K", "283.15K", 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine.Set("celsius", tt.set)
			w := tempconv.Celsius(tt.want)
			if *celsius != w {
				t.Errorf("celsius got %s, want %s", celsius.String(), w.String())
			}
		})
	}
}

func Test_Kelvin(t *testing.T) {
	tests := []struct {
		name string
		set  string
		want float64
	}{
		{"default", "", 20.0},
		{"30K", "30K", 30.0},
		{"30C", "30C", 303.15},
		{"40°C", "40°C", 313.15},
		{"50F", "50F", 283.15},
		{"59°F", "59°F", 288.15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine.Set("kelvin", tt.set)
			w := tempconv.Kelvin(tt.want)
			if *kelvin != w {
				t.Errorf("kelvin got %s, want %s", kelvin.String(), w.String())
			}
		})
	}
}
