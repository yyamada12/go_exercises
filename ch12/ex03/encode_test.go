package sexpr

import (
	"bytes"
	"fmt"
	"testing"
)

type stringerImpl struct {
	s string
}

func (s stringerImpl) String() string {
	return s.s
}

func Test_encode(t *testing.T) {
	var i interface{}
	var s fmt.Stringer = stringerImpl{"string"}

	tests := []struct {
		name string
		v    interface{}
		want []byte
	}{
		{"boolean true", true, []byte("t")},
		{"boolean false", false, []byte("nil")},
		{"float32", 3.2, []byte(fmt.Sprintf("%f", float32(3.2)))},
		{"float64", 64.646464646464, []byte(fmt.Sprintf("%f", float64(64.646464646464)))},
		{"complex64", 3.2 + 3.2i, []byte(fmt.Sprintf("#C(%f %f)", float32(3.2), float32(3.2)))},
		{"complex128", 6.4 + 64.646464646464i, []byte(fmt.Sprintf("#C(%f %f)", float64(6.4), float64(64.646464646464)))},
		{"interface nil", &i, []byte("nil")},
		{"interface Writer", &s, []byte("(sexpr.stringerImpl ((s \"string\")))")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.v)
			if err != nil {
				t.Errorf("Marshal got err: %s", err)
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Marshal(%v) got %s want %s", tt.v, got, tt.want)
			}
		})
	}
}
