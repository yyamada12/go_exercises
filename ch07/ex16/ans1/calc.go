package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	prv    float64
	prvOpe string
	crtNum string
)

func main() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/num", num)
	http.HandleFunc("/ope", ope)
	http.HandleFunc("/eval", eval)
	http.HandleFunc("/clear", clear)
	fmt.Println("listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func topHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, 0)
}

func num(w http.ResponseWriter, req *http.Request) {
	n := req.URL.Query().Get("num")
	if n == "." && crtNum == "" {
		crtNum = "0."
	} else if !(n == "0" && crtNum == "0") {
		crtNum += n
	}

	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, crtNum)
}

func ope(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))

	if crtNum == "" {
		prvOpe = req.URL.Query().Get("ope")
		t.Execute(w, prv)
		return
	}

	crt, _ := strconv.ParseFloat(crtNum, 64)
	crtNum = ""
	if prvOpe != "" {
		switch prvOpe {
		case "add":
			prv += crt
		case "sub":
			prv -= crt
		case "mul":
			prv *= crt
		case "div":
			prv /= crt
		}
	} else {
		prv = crt
	}
	prvOpe = req.URL.Query().Get("ope")
	t.Execute(w, prv)
}

func eval(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))

	if crtNum == "" || prvOpe == "" {
		t.Execute(w, prv)
		return
	}

	crt, _ := strconv.ParseFloat(crtNum, 64)
	crtNum = ""

	switch prvOpe {
	case "add":
		prv += crt
	case "sub":
		prv -= crt
	case "mul":
		prv *= crt
	case "div":
		prv /= crt
	}
	prvOpe = ""

	t.Execute(w, prv)
}

func clear(w http.ResponseWriter, req *http.Request) {
	crtNum = ""
	prvOpe = ""
	prv = 0

	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, prv)
}
