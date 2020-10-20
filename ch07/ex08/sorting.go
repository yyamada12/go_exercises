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
	var x conditionMemorySort

	fmt.Println("\nbyYear:")
	x = addSortCondition(x, byYear, true)
	sort.Sort(x)
	printTracks(tracks)

	fmt.Println("byTitle:")
	x = addSortCondition(x, byTitle, false)
	sort.Sort(x)
	printTracks(tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

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

func addSortCondition(x conditionMemorySort, s sortCondition, reverse bool) conditionMemorySort {
	switch s {
	case byTitle:
		if reverse {
			return conditionMemorySort{sort.Reverse(customSort{tracks, lessByTitle}), &x}
		}
		return conditionMemorySort{customSort{tracks, lessByTitle}, &x}
	case byArtist:
		if reverse {
			return conditionMemorySort{sort.Reverse(customSort{tracks, lessByArtist}), &x}
		}
		return conditionMemorySort{customSort{tracks, lessByArtist}, &x}
	case byAlbum:
		if reverse {
			return conditionMemorySort{sort.Reverse(customSort{tracks, lessByAlbum}), &x}
		}
		return conditionMemorySort{customSort{tracks, lessByAlbum}, &x}
	case byYear:
		if reverse {
			return conditionMemorySort{sort.Reverse(customSort{tracks, lessByYear}), &x}
		}
		return conditionMemorySort{customSort{tracks, lessByYear}, &x}
	case byLength:
		if reverse {
			return conditionMemorySort{sort.Reverse(customSort{tracks, lessByLength}), &x}
		}
		return conditionMemorySort{customSort{tracks, lessByLength}, &x}
	default:
		return x
	}
}

type conditionMemorySort struct {
	crt sort.Interface
	prv *conditionMemorySort
}

func (x conditionMemorySort) Len() int      { return x.crt.Len() }
func (x conditionMemorySort) Swap(i, j int) { x.crt.Swap(i, j) }
func (x conditionMemorySort) Less(i, j int) bool {
	if !x.crt.Less(i, j) && !x.crt.Less(j, i) && x.prv.crt != nil {
		return (*x.prv).Less(i, j)
	}
	return x.crt.Less(i, j)
}
