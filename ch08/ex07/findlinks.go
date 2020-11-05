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

	"github.com/yyamada12/go_exercises/ch08/ex07/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	for _, targetURL := range os.Args[1:] {
		u, err := url.Parse(targetURL)
		if err != nil {
			log.Print(err)
			return
		}
		targetHostnames = append(targetHostnames, u.Hostname())
	}

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var targetHostnames []string

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token

	// Download page
	if err := download(url); err != nil {
		log.Print(err)
	}

	// Extract link
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	<-tokens // release the token
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

	// replace url
	r := strings.NewReplacer(
		"https://"+hostname+"/", "/",
		"https://"+hostname, "/",
		"http://"+hostname+"/", "/",
		"http://"+hostname, "/",
	)
	data = []byte(r.Replace(string(data)))

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
