// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

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

func main() {
	http.HandleFunc("/tracks", topHandler)
	fmt.Println("listening at http://localhost:8000/tracks")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

var prvCond string

func topHandler(w http.ResponseWriter, r *http.Request) {
	cond := r.URL.Query().Get("sortBy")
	reverse := prvCond == cond
	prvCond = cond
	sortTracksBy(cond, reverse)
	printTracks(tracks, w)
}

func printTracks(tracks []*Track, w io.Writer) {
	t := template.Must(template.ParseFiles("template/index.html"))
	err := t.Execute(w, tracks)
	if err != nil {
		log.Fatal(err)
	}
}

func sortTracksBy(cond string, reverse bool) {
	switch cond {
	case "Title":
		if reverse {
			sort.Stable(sort.Reverse(customSort{tracks, lessByTitle}))
		} else {
			sort.Stable(customSort{tracks, lessByTitle})
		}
	case "Artist":
		if reverse {
			sort.Stable(sort.Reverse(customSort{tracks, lessByArtist}))
		} else {
			sort.Stable(customSort{tracks, lessByArtist})
		}
	case "Album":
		if reverse {
			sort.Stable(sort.Reverse(customSort{tracks, lessByAlbum}))
		} else {
			sort.Stable(customSort{tracks, lessByAlbum})
		}
	case "Year":
		if reverse {
			sort.Stable(sort.Reverse(customSort{tracks, lessByYear}))
		} else {
			sort.Stable(customSort{tracks, lessByYear})
		}
	case "Length":
		if reverse {
			sort.Stable(sort.Reverse(customSort{tracks, lessByLength}))
		} else {
			sort.Stable(customSort{tracks, lessByLength})
		}
	}
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
