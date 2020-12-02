// Http is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{m: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	m map[string]dollars
	sync.RWMutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	type Item struct {
		Item  string
		Price dollars
	}
	t := template.Must(template.ParseFiles("template/index.html"))

	db.RLock()
	defer db.RUnlock()
	var items []Item
	for item, price := range db.m {
		items = append(items, Item{item, price})
	}
	err := t.Execute(w, items)
	if err != nil {
		log.Print(err)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item required\n")
		return
	}

	db.RLock()
	defer db.RUnlock()
	if price, ok := db.m[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item required\n")
		return
	}

	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price required\n")
		return
	}

	p, err := parsePrice(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}

	db.Lock()
	defer db.Unlock()
	if _, ok := db.m[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item already exist: %s\n", item)
		return
	}

	db.m[item] = p
	fmt.Fprintf(w, "%s: %s created\n", item, p)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item required\n")
		return
	}

	price := req.URL.Query().Get("price")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price required\n")
		return
	}

	p, err := parsePrice(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %s\n", price)
		return
	}

	db.Lock()
	defer db.Unlock()
	if _, ok := db.m[item]; !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item not exists: %s\n", item)
	}

	db.m[item] = p
	fmt.Fprintf(w, "price of %s updated for %s\n", item, p)
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item required\n")
		return
	}
	db.Lock()
	defer db.Unlock()
	if _, ok := db.m[item]; !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item not exists: %s\n", item)
	}
	delete(db.m, item)
	fmt.Fprintf(w, "%s is deleted\n", item)
}

func parsePrice(p string) (dollars, error) {
	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return 0, err
	}
	if price < 0 {
		return 0, fmt.Errorf("price cannot be negative value: %f", price)
	}
	return dollars(price), nil
}
