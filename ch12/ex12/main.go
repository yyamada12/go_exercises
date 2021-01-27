package main

import (
	"fmt"
	"log"
	"net/http"

	params "github.com/yyamada12/go_exercises/ch12/ex12/params"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels           []string `http:"l"`
		MaxResults       int      `http:"max"`
		Exact            bool     `http:"x"`
		Email            string   `valid:"email"`
		CreditCardNumber string   `http:"credit" valid:"credit_card"`
		Zip              int      `valid:"post_code"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
$ go build -o search
$ ./search &
$ ./fetch 'http://localhost:12345/search'
Search: {Labels:[] MaxResults:10 Exact:false Email: CreditCardNumber: Zip:0}
$ ./fetch 'http://localhost:12345/search?email=foo'
foo is invalid email address
$ ./fetch 'http://localhost:12345/search?email=foo@gmail.com'
Search: {Labels:[] MaxResults:10 Exact:false Email:foo@gmail.com CreditCardNumber: Zip:0}
$ ./fetch 'http://localhost:12345/search?credit=1111-2222-3333-4444'
1111-2222-3333-4444 is invalid credit card number
$ ./fetch 'http://localhost:12345/search?credit=4929-5678-9012-3456'
Search: {Labels:[] MaxResults:10 Exact:false Email: CreditCardNumber:4929-5678-9012-3456 Zip:0}
$ ./fetch 'http://localhost:12345/search?zip=aa'
invalid postal code format
$ ./fetch 'http://localhost:12345/search?zip=1'
postal code cannot be shorter than 2 characters
$ ./fetch 'http://localhost:12345/search?zip=10001'
Search: {Labels:[] MaxResults:10 Exact:false Email: CreditCardNumber: Zip:10001}
*/
