package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var h = template.Must(template.New("template").Parse(`
		<h1>data</h1>
		<table>
		<tr style='text-align: left'>
			<th>Name</th>
			<th>Price</th>
		</tr>
		{{range $k, $v := .}}
		<tr>
			<td>{{$k}}</td>
			<td>{{$v}}</td>
		</tr>
		{{end}}
		</table>
	`))
	if err := h.Execute(w, db); err != nil {
		// ...
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
