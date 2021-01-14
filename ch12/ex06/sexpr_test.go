package sexpr

import (
	"reflect"
	"testing"

	"gopl.io/ch12/sexpr"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "",
		Year:     0,
		// Actor:    map[string]string{
		// "Dr. Strangelove":            "Peter Sellers",
		// "Grp. Capt. Lionel Mandrake": "Peter Sellers",
		// "Pres. Merkin Muffley":       "Peter Sellers",
		// "Gen. Buck Turgidson":        "George C. Scott",
		// "Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		// `Maj. T.J. "King" Kong`:      "Slim Pickens",
		// },
		// Oscars: []string{
		// 	"Best Actor (Nomin.)",
		// 	"Best Adapted Screenplay (Nomin.)",
		// 	"Best Director (Nomin.)",
		// 	"Best Picture (Nomin.)",
		// },
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := sexpr.Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}
