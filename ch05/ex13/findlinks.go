// Findlinks crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/yyamada12/go_exercises/ch05/ex13/links"
)

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

var targetHostnames []string

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	for _, item := range worklist {
		u, err := url.Parse(item)
		if err != nil {
			log.Print(err)
			return
		}
		targetHostnames = append(targetHostnames, u.Hostname())
	}

	seen := make(map[string]bool)
	i := 0
	for len(worklist) > 0 {
		i++
		if i > 10 {
			break
		}
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(u string) []string {
	fmt.Println(u)

	// Extract link
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}

	// get hostname from u
	uu, err := url.Parse(u)
	if err != nil {
		log.Print(err)
		return list
	}
	hostname := uu.Hostname()
	urlpath := uu.Path
	if urlpath == "" {
		urlpath = "/index.html"
	} else if strings.HasSuffix(urlpath, "/") {
		urlpath += "index.html"
	} else if !strings.Contains(path.Base(urlpath), ".") {
		urlpath += "/index.html"
	}

	if contains(targetHostnames, hostname) {
		if err := os.MkdirAll(path.Dir(hostname+urlpath), 0755); err != nil {
			log.Print(err)
			return list
		}
		if err := fetch(hostname+urlpath, u); err != nil {
			log.Print(err)
		}
	}

	return list
}

func contains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func fetch(fileName, u string) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error get %s: %s", u, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, data, 0644)
	return err
}
