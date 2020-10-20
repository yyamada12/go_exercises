package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"testing"
	"text/tabwriter"
)

func Test_sorting(t *testing.T) {
	// setup
	tracks = []*Track{
		{"A", "a", "", 2000, length("1s")},
		{"A", "a", "", 1999, length("1s")},
		{"A", "b", "", 1999, length("1s")},
		{"A", "b", "", 2000, length("1s")},
		{"B", "a", "", 2000, length("1s")},
		{"B", "a", "", 1999, length("1s")},
		{"B", "b", "", 1999, length("1s")},
		{"B", "b", "", 2000, length("1s")},
	}
	tracks2 := []*Track{
		{"A", "a", "", 2000, length("1s")},
		{"B", "a", "", 1999, length("1s")},
		{"A", "b", "", 1999, length("1s")},
		{"B", "a", "", 2000, length("1s")},
		{"A", "a", "", 1999, length("1s")},
		{"B", "b", "", 1999, length("1s")},
		{"B", "b", "", 2000, length("1s")},
		{"A", "b", "", 2000, length("1s")},
	}

	// execute

	// sort tracks by conditionMemorySort
	var x conditionMemorySort
	x = addSortCondition(x, byYear, true)
	sort.Sort(x)
	x = addSortCondition(x, byArtist, false)
	sort.Sort(x)
	x = addSortCondition(x, byTitle, false)
	sort.Sort(x)

	// sort tracks2 by stableSort
	sort.Stable(sort.Reverse(customSort{tracks2, lessByYear}))
	sort.Stable(customSort{tracks2, lessByArtist})
	sort.Stable(customSort{tracks2, lessByTitle})

	// assert
	if !reflect.DeepEqual(tracks, tracks2) {
		t.Errorf("result of conditionMemorySort not equal to result of Stable sort.")
		t.Errorf("got:\n%s", printTracksString(tracks))
		t.Errorf("want:\n%s", printTracksString(tracks2))
	}

}

func printTracksString(tracks []*Track) string {
	buf := new(bytes.Buffer)
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
	return buf.String()
}
