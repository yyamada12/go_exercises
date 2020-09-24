// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yyamada12/go_exercises/ch04/ex11/editor"
	"github.com/yyamada12/go_exercises/ch04/ex11/github"
)

const titleText = "---------------------- Issue Title ----------------------\n"
const bodyText = "---------------------- Issue Body ----------------------\n"

func main() {
	if len(os.Args) < 4 {
		printUsage()
		return
	}
	cmd, owner, repo := os.Args[1], os.Args[2], os.Args[3]
	token, editorName := os.Getenv("TOKEN"), os.Getenv("EDITOR")
	// default editor
	if editorName == "" {
		editorName = "vim"
	}
	// create new github operator
	operator := github.NewIssueOperator(token, owner, repo)
	switch cmd {
	case "get":
		if len(os.Args) < 5 {
			printGetUsage()
			return
		}
		number := os.Args[4]
		res, err := operator.GetIssue(number)
		handleErr(err)
		fmt.Println("issue created!")
		printIssue(res)
	case "create":
		if token == "" {
			printCreateUsage()
			return
		}
		// input title and body
		title, body, err := inputTitleAndBodyWithEditor(editorName)
		handleErr(err)
		// create issue
		res, err := operator.CreateIssue(title, body)
		handleErr(err)
		printIssue(res)
	case "update":
		if len(os.Args) < 5 || token == "" {
			printUpdateUsage()
			return
		}
		number := os.Args[4]
		// get current issue
		res, err := operator.GetIssue(number)
		handleErr(err)
		// modify issue
		title, body, err := modifyTitleAndBodyWithEditor(editorName, res.Title, res.Body)
		handleErr(err)
		// update issue
		res, err = operator.UpdateIssue(number, title, body)
		handleErr(err)
		fmt.Println("issue updated!")
		printIssue(res)
	case "close":
		if len(os.Args) < 5 || token == "" {
			printCloseUsage()
			return
		}
		number := os.Args[4]
		res, err := operator.CloseIssue(number)
		handleErr(err)
		fmt.Println("issue closed!")
		printIssue(res)
	default:
		printUsage()
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printIssue(res *github.IssueResult) {
	fmt.Printf("#%-5d %9.9s %.55s\n%s\n",
		res.Number, res.User.Login, res.Title, res.Body)
}

func inputTitleAndBodyWithEditor(editorName string) (title, body string, err error) {
	text, err := editor.ReadInputWithEditor(editorName, titleText+"\n"+bodyText+"\n", ".issue")
	if err != nil {
		return
	}
	return parseTitleAndBody(text)
}

func modifyTitleAndBodyWithEditor(editorName, oldTitle, oldBody string) (title, body string, err error) {
	text, err := editor.ReadInputWithEditor(editorName, titleText+oldTitle+"\n"+bodyText+oldBody, ".issue")
	if err != nil {
		return
	}
	return parseTitleAndBody(text)
}

func parseTitleAndBody(text string) (title, body string, err error) {
	text = strings.Replace(text, titleText, "", 1)
	bodyAndTitle := strings.Split(text, bodyText)
	if len(bodyAndTitle) != 2 {
		err = errors.New("parse text error: Do not edit header line \"---------------------- Issue xxxx ----------------------\"")
		return
	}
	return strings.TrimSpace(bodyAndTitle[0]), strings.TrimSpace(bodyAndTitle[1]), nil
}

func printUsage() {
	fmt.Println(`USAGE:
	TOKEN=xxxx EDITOR=xxxx go run issue.go [COMMAND] owner repo

TOKEN:
	please set to ENV variable your github personal access token
EDITOR:
	please set to ENV variable your favorite editor for edit issue	(default vim)
COMMAND:
	get [issue number]		get issue by issue number
	create					create new issue 						(need TOKEN)
	update [issue number]	update issue by issue number 			(need TOKEN)
	close [issue number]	close issue by issue number 			(need TOKEN)`)
}

func printGetUsage() {
	fmt.Println(`USAGE:
	go run issue.go get owner repo issue_number`)
}

func printCreateUsage() {
	fmt.Println(`USAGE:
	TOKEN=xxxx go run issue.go create owner repo`)
}

func printUpdateUsage() {
	fmt.Println(`USAGE:
	TOKEN=xxxx go run issue.go update owner repo issue_number`)
}

func printCloseUsage() {
	fmt.Println(`USAGE:
	TOKEN=xxxx go run issue.go close owner repo issue_number`)
}
