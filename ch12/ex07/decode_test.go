package sexpr

import (
	"bytes"
	"reflect"
	"testing"

	"gopl.io/ch12/sexpr"
)

func TestDecoder(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove1 := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	strangelove2 := Movie{
		Title:    "Dr. Strangelove2",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb2",
		Year:     1965,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
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
