package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a) // "[2 3 4 5 0 1]"

}

func rotate(s []int, n int) []int {
	return swapRotate(s, 0, len(s), n)
}

// A B -> B A
// A: [start, start+1, ... , start+n-1]
// B: [start+n, start+n+1, ... , start+len-1]
func swapRotate(s []int, start, len, n int) []int {
	if n == 0 || n == len {
		return s
	}

	// len(A) == len(B)
	if n == len-n {
		return blockSwap(s, start, start+n, n)
	}

	if n < len-n {
		// if len(A) < len(B)
		// A BL BR -> BR BL A
		s = blockSwap(s, start, start+len-n, n)
		// BR BL A -> BL BR A (= B A)
		return swapRotate(s, start, len-n, n)
	} else {
		// if len(A) > len(B)
		// AL AR B -> B AR AL
		s = blockSwap(s, start, n, len-n)
		// B AR AL -> B AL AR (= B A)
		return swapRotate(s, start+len-n, n-(len-n), n)
	}
}

// swap [l, l+1, ... , l+n] and [r, r+1, ... , r+n]
func blockSwap(s []int, l, r, n int) []int {
	for i := 0; i < n; i++ {
		s[l+i], s[r+i] = s[r+i], s[l+i]
	}
	return s
}
