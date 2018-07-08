package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
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

const templ = `
<h1>Tracks</h1>
<table>
<tr style='text-align:left'>
	<th onclick="clickTable(this, 'title')">Title</th>
	<th onclick="clickTable(this, 'artist')">Artist</th>
	<th onclick="clickTable(this, 'album')">Album</th>
	<th onclick="clickTable(this, 'year')">Year</th>
	<th onclick="clickTable(this, 'length')">Length</th>
</tr>
{{range .}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`

func htmlTracks(w io.Writer, tracks []*Track) {
	var h = template.Must(template.New("track").Parse(templ))
	if err := h.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}

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
		} else if s == "album" && x.t[i].Album != x.t[j].Album {
			return x.t[i].Album < x.t[j].Album
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

// Tracks Tracks
type Tracks []*Track

// ServeHTTP serve the http request
func (t Tracks) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// parse form to get which column should be sorted by
	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// / to get the main page, /sort to get the sorted table
	switch req.URL.Path {
	case "/":
		b, _ := ioutil.ReadFile("a.html")
		fmt.Fprintln(w, string(b))
	case "/sort":
		if len(req.Form["q"]) > 0 {
			query := req.Form["q"][0]
			querys := strings.Split(query, ",")
			// sort by given columns
			sort.Sort(multiSort{t, querys})
		}
		htmlTracks(w, t)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	log.Fatal(http.ListenAndServe("localhost:8000", Tracks(tracks)))
}
