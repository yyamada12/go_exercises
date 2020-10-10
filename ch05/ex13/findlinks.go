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
	for len(worklist) > 0 {
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

func crawl(url string) []string {
	fmt.Println(url)

	// Download page
	if err := download(url); err != nil {
		log.Print(err)
	}

	// Extract link
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
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

func download(url string) error {
	hostname, urlpath, err := parseURL(url)
	if err != nil {
		return err
	}

	// download target host only
	if !contains(targetHostnames, hostname) {
		return nil
	}

	// fetch data
	data, err := fetch(url)
	if err != nil {
		return err
	}

	// save data
	filepath := hostname + urlpath
	err = os.MkdirAll(path.Dir(filepath), 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, data, 0644)
	return err
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error get %s: %s", url, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func parseURL(u string) (string, string, error) {
	uu, err := url.Parse(u)
	if err != nil {
		return "", "", err
	}
	hostname := uu.Hostname()
	urlpath := uu.Path
	if strings.HasSuffix(urlpath, "/") {
		urlpath += "index.html"
	} else if urlpath == "" || !strings.Contains(path.Base(urlpath), ".") {
		urlpath += "/index.html"
	}
	return hostname, urlpath, nil
}
