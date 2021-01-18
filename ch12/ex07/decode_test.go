package sexpr

import (
	"bytes"
	"reflect"
	"testing"

	"gopl.io/ch12/sexpr"
)

func TestDecoder(t *testing.T) {
	type Movie struct {
		Title string
	}
	strangelove1 := Movie{
		Title: "Dr. Strangelove",
	}
	strangelove2 := Movie{
		Title: "Dr. Strangelove2",
	}

	// Encode it
	data1, _ := sexpr.Marshal(strangelove1)
	data2, _ := sexpr.Marshal(strangelove2)
	data := append(data1, data2...)

	dec := NewDecoder(bytes.NewReader(data))
	var movie1, movie2 Movie
	dec.Decode(&movie1)
	dec.Decode(&movie2)
	if !reflect.DeepEqual(strangelove1, movie1) {
		t.Errorf("first Decode(data) got %v, want %v", strangelove1, movie1)
	}
	if !reflect.DeepEqual(strangelove1, movie1) {
		t.Errorf("second Decode(data) got %v, want %v", strangelove2, movie2)
	}
}
