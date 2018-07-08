package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// protect database
var mu sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	mu.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	mu.Unlock()
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "parse error")
		return
	}
	mu.Lock()
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item already in database\n")
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "%s added to database\n", item)
	}
	mu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "parse error")
		return
	}
	mu.Lock()
	if _, ok := db[item]; ok {
		db[item] = dollars(price)
		fmt.Fprintf(w, "%s updated\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not in database\n")
	}
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%s deleted\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not in database\n")
	}
	mu.Unlock()
}
