// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type sortCondition int

const (
	_ sortCondition = iota
	byTitle
	byArtist
	byAlbum
	byYear
	byLength
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	x := conditionMemorySort{tracks, func(x, y *Track) bool { return false }, nil}
	fmt.Println("byTitle:")
	x = addSortCondition(x, byTitle)
	sort.Sort(x)
	printTracks(tracks)
	fmt.Println("\nbyYear:")
	x = addSortCondition(x, byYear)
	sort.Sort(x)
	printTracks(tracks)

}

func lessByTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func lessByArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func lessByAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func lessByYear(x, y *Track) bool {
	return x.Year < y.Year
}

func lessByLength(x, y *Track) bool {
	return x.Length < y.Length
}

func addSortCondition(prev conditionMemorySort, s sortCondition) conditionMemorySort {
	switch s {
	case byTitle:
		return conditionMemorySort{prev.t, lessByTitle, &prev}
	case byArtist:
		return conditionMemorySort{prev.t, lessByArtist, &prev}
	case byAlbum:
		return conditionMemorySort{prev.t, lessByAlbum, &prev}
	case byYear:
		return conditionMemorySort{prev.t, lessByYear, &prev}
	case byLength:
		return conditionMemorySort{prev.t, lessByLength, &prev}
	default:
		return prev
	}
}

type conditionMemorySort struct {
	t    []*Track
	less func(x, y *Track) bool
	prev *conditionMemorySort
}

func (x conditionMemorySort) Len() int      { return len(x.t) }
func (x conditionMemorySort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x conditionMemorySort) Less(i, j int) bool {
	if !x.less(x.t[i], x.t[j]) && !x.less(x.t[j], x.t[i]) && x.prev != nil {
		return x.prev.Less(i, j)
	}
	return x.less(x.t[i], x.t[j])
}
