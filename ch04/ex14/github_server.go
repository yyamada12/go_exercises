package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/yyamada12/go_exercises/ch04/ex14/github"
)

var issueSearchResult *github.IssuesSearchResult

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	var err error
	issueSearchResult, err = github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", topHandler)
	http.HandleFunc("/bugreports", bugreportHandler)
	http.HandleFunc("/milestones", milestoneHandler)
	http.HandleFunc("/users", userHandler)
	fmt.Println("listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func printUsage() {
	fmt.Println(`USAGE:
	go run github_server.go [SEARCH QUERIES]
	
	use SEARCH QUERIES github format like "repo:golang/go is:open"
	`)
}

func topHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/index.html"))
	err := t.Execute(w, issueSearchResult)
	if err != nil {
		log.Fatal(err)
	}
}

func bugreportHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/bugreport.html"))
	bugreports := extractBugreport(issueSearchResult)
	err := t.Execute(w, bugreports)
	if err != nil {
		log.Fatal(err)
	}
}

func milestoneHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/milestone.html"))
	milestones := extractMilestone(issueSearchResult)
	err := t.Execute(w, milestones)
	if err != nil {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/user.html"))
	users := extractUser(issueSearchResult)
	err := t.Execute(w, users)
	if err != nil {
		log.Fatal(err)
	}
}

func extractBugreport(result *github.IssuesSearchResult) []*github.Issue {
	res := []*github.Issue{}
	for _, issue := range result.Items {
		for _, l := range issue.Labels {
			if l.Name == "bug" {
				res = append(res, issue)
			}
		}
	}
	return res
}

func extractMilestone(result *github.IssuesSearchResult) []*github.Milestone {
	m := make(map[string]*github.Milestone)
	for _, issue := range result.Items {
		if issue.Milestone != nil {
			m[issue.Milestone.HTMLURL] = issue.Milestone
		}
	}
	res := []*github.Milestone{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func extractUser(result *github.IssuesSearchResult) []*github.User {
	m := make(map[string]*github.User)
	for _, issue := range result.Items {
		m[issue.User.HTMLURL] = issue.User
		for _, asignee := range issue.Assignees {
			m[asignee.HTMLURL] = asignee
		}
	}
	res := []*github.User{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
