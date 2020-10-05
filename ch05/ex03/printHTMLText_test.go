// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"log"
	"os"
	"reflect"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func Test_visit(t *testing.T) {
	// setup
	fp, err := os.Open("sample.html")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	doc, err := html.Parse(fp)
	if err != nil {
		log.Fatal(err)
	}

	// execute
	got := visit(nil, doc)

	// assert
	want := []string{"The Go Programming Language", "Black Lives Matter.", "Support the Equal Justice Initiative.", "Documents", "Packages", "The Project", "Help", "Blog", "Play", "Search", "Go is an open source programming language that makes it easy to", "build", "simple", ",", "reliable", ", and", "efficient", "software.", "Download Go", "Binary distributions available for", "Linux, macOS, Windows, and more.", "Try Go", "Open in Playground", "// You can edit this code!", "// Click here and start typing.", "package main", "import \"fmt\"", "func main() {", "fmt.Println(\"Hello, 世界\")", "}", "Hello, 世界", "Hello, World!", "Conway's Game of Life", "Fibonacci Closure", "Peano Integers", "Concurrent pi", "Concurrent Prime Sieve", "Peg Solitaire Solver", "Tree Comparison", "Run", "Share", "Tour", "Featured articles", "Read more >", "Featured video", "Copyright", "Terms of Service", "Privacy Policy", "Report a website issue", "Supported by Google"}

	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
