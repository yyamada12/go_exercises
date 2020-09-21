// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

/*
$ ./issues repo:golang/go is:open json decoder
*/

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	lastMonth, lastYear, past := separateByCreatedAt(result.Items)

	fmt.Println()
	fmt.Println("Issues created last month")
	printIssues(lastMonth)
	fmt.Println()
	fmt.Println("Issues created last year")
	printIssues(lastYear)
	fmt.Println()
	fmt.Println("Issues created more than a year ago")
	printIssues(past)
}

func printIssues(items []*github.Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func separateByCreatedAt(items []*github.Issue) (lastMonth, lastYear, past []*github.Issue) {
	for _, item := range items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			lastMonth = append(lastMonth, item)
		} else if item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			lastYear = append(lastYear, item)
		} else {
			past = append(past, item)
		}

	}
	return
}
