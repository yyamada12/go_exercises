package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/yyamada12/go_exercises/ch07/ex16/ans2/eval"
)

func main() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/calc", calc)
	fmt.Println("listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func topHandler(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, "")
}

func calc(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))

	// read formula from stdin
	text := req.URL.Query().Get("formula")

	// parse formula
	expr, err := eval.Parse(text)
	if err != nil {
		log.Print(err)
		t.Execute(w, "wrong formula")
		return
	}

	// check and extract vars
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		log.Print(err)
		t.Execute(w, "wrong formula")
		return
	}

	// eval formula
	val := expr.Eval(eval.Env{})
	t.Execute(w, val)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
