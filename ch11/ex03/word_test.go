// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

// randomNonPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	var fst, lst rune
	for !unicode.IsLetter(fst) {
		fst = rune(rng.Intn(0x1000))
	}

	for !unicode.IsLetter(lst) || unicode.ToLower(fst) == unicode.ToLower(lst) {
		lst = rune(rng.Intn(0x1000))
	}

	return string(fst) + string(runes) + string(lst)
}

func TestRandomnonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
