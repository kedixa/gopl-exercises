package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track track
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

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (x byLength) Len() int           { return len(x) }
func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLength) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type multiSort struct {
	t     []*Track
	state []string
}

func (x multiSort) Len() int { return len(x.t) }
func (x multiSort) Less(i, j int) bool {
	// for each state
	for k := len(x.state) - 1; k >= 0; k-- {
		s := x.state[k]
		if s == "title" && x.t[i].Title != x.t[j].Title {
			return x.t[i].Title < x.t[j].Title
		} else if s == "artist" && x.t[i].Artist != x.t[j].Artist {
			return x.t[i].Artist < x.t[j].Artist
		} else if s == "year" && x.t[i].Year != x.t[j].Year {
			return x.t[i].Year < x.t[j].Year
		} else if s == "length" && x.t[i].Length != x.t[j].Length {
			return x.t[i].Length < x.t[j].Length
		}
	}
	// default compare
	return x.t[i].Title < x.t[j].Title
}
func (x multiSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println("\nmultiSort:")
	// Attention: check title first, then year, not the opposite
	sort.Sort(multiSort{tracks, []string{"year", "title"}})
	printTracks(tracks)
	// disrupt the order
	sort.Sort(sort.Reverse(byLength(tracks)))

	fmt.Println("\nStable Sort:")
	sort.Stable(byYear(tracks))
	sort.Stable(byTitle(tracks))
	printTracks(tracks)
	// The two results are the same
}
