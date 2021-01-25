// Package params provides a reflection-based parser for URL parameters.
package params

import "testing"

func TestPack(t *testing.T) {
	type data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool
	}

	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"int param with default bool param", data{MaxResults: 1}, "?max=1&exact=false"},
		{"slice and bool param with default int param", data{Labels: []string{"label1", "label2"}, Exact: true}, "?l=label1&l=label2&max=0&exact=true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pack(tt.v)
			if err != nil {
				t.Errorf("Pack(%v) got error: %s", tt.v, err)
			}
			if got != tt.want {
				t.Errorf("Pack(%v) = %q, want %q", tt.v, got, tt.want)
			}
		})
	}
}
