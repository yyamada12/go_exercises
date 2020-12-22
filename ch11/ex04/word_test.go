package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindromeWithPunctuationSpace(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	if n == 0 {
		return ""
	}
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	marks := `.?!,:;-[]{}()'" `
	mark := marks[rng.Intn(len(marks))]

	m := rng.Intn(n)
	return string(runes[:m]) + string(mark) + string(runes[m:])
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindromeWithPunctuationSpace(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
