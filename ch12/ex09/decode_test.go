package sexpr

import (
	"bytes"
	"io"
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
	strangelove := Movie{
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

	// Encode it
	data, _ := sexpr.Marshal(strangelove)

	dec := NewDecoder(bytes.NewReader(data))
	for {
		tok, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Errorf("token err: %s", err)
		}
		switch v := tok.(type) {
		case Symbol:
			t.Logf(string(v))
		case String:
			t.Logf(string(v))
		case Int:
			t.Logf("%d", int(v))
		case StartList:
			t.Logf("(")
		case EndList:
			t.Logf(")")
		}
	}

}
