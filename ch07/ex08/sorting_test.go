package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"text/tabwriter"
	"time"
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

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[int(rand.Int63()%int64(len(letters)))]
	}
	return string(b)
}

func createTracks(n int) []*Track {
	tracks := make([]*Track, n)
	for i, _ := range tracks {
		tracks[i] = &Track{randString(5), randString(5), randString(5), rand.Int() % 3000, length(fmt.Sprintf("%dm%ds", rand.Int()%60, rand.Int()%60))}
	}
	return tracks
}

func sortByConditionMemorySort(n int) {
	var x conditionMemorySort
	for j := 0; j < n; j++ {
		switch rand.Int63() % 5 {
		case 0:
			x = addSortCondition(x, byTitle, false)
			sort.Sort(x)
		case 1:
			x = addSortCondition(x, byArtist, false)
			sort.Sort(x)
		case 2:
			x = addSortCondition(x, byAlbum, false)
			sort.Sort(x)
		case 3:
			x = addSortCondition(x, byYear, false)
			sort.Sort(x)
		case 4:
			x = addSortCondition(x, byLength, false)
			sort.Sort(x)
		}
	}
}

func sortByStableSort(n int) {
	for j := 0; j < n; j++ {
		switch rand.Int63() % 5 {
		case 0:
			sort.Stable(customSort{tracks, lessByTitle})
		case 1:
			sort.Stable(customSort{tracks, lessByArtist})
		case 2:
			sort.Stable(customSort{tracks, lessByAlbum})
		case 3:
			sort.Stable(customSort{tracks, lessByYear})
		case 4:
			sort.Stable(customSort{tracks, lessByLength})
		}
	}
}

func BenchmarkConditionMemorySort_10times_100tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(100)
	for i := 0; i < b.N; i++ {
		sortByConditionMemorySort(10)
	}
}

func BenchmarkStableSort_10times_100tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(100)
	for i := 0; i < b.N; i++ {
		sortByStableSort(10)
	}
}

func BenchmarkConditionMemorySort_100times_100tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(100)
	for i := 0; i < b.N; i++ {
		sortByConditionMemorySort(100)
	}
}

func BenchmarkStableSort_100times_100tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(100)
	for i := 0; i < b.N; i++ {
		sortByStableSort(100)
	}
}

func BenchmarkConditionMemorySort_10times_10000tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(10000)
	for i := 0; i < b.N; i++ {
		sortByConditionMemorySort(10)
	}
}

func BenchmarkStableSort_10times_10000tracks(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	tracks = createTracks(10000)
	for i := 0; i < b.N; i++ {
		sortByStableSort(10)
	}
}
