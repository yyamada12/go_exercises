package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const indexFileName = "xkcd_%d.json"
const indexFileDir = "index"
const comicNum = 2363
const xkcdURL = "http://xkcd.com/%d/info.0.json"

type comic struct {
	Num        int
	Day        string
	Month      string
	Year       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
}

func main() {

	err := createIndexIfNotExit()
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	searchAllComics(os.Args[1])
}

func printUsage() {
	fmt.Println(`USAGE:
	go run xkcd.go TERM`)
}

func createIndexIfNotExit() error {
	// if directory is not exist
	if _, err := os.Stat(indexFileDir); err != nil {
		// create directory for index
		if err = os.Mkdir(indexFileDir, 0755); err != nil {
			return err
		}
	}
	// if counts of index file is less than comicNum, create index
	if files, err := ioutil.ReadDir(indexFileDir); err != nil {
		return err
	} else if len(files) < comicNum-1 { // -1 because comic number 404 is not available
		createIndex()
	}
	return nil
}

func createIndex() {
	var wg sync.WaitGroup
	for i := 1; i <= comicNum; i++ {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			filepath := fmt.Sprintf(indexFileDir+"/"+indexFileName, number)
			// if file not yet created
			if _, err := os.Stat(filepath); err != nil {
				if err := saveAllComics(number); err != nil {
					log.Println(err)
				}
			}
		}(i)
	}
	wg.Wait()
}

func saveAllComics(number int) error {
	comicURL := fmt.Sprintf(xkcdURL, number)
	resp, err := http.Get(comicURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error get comic %d: %s", number, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf(indexFileDir+"/"+indexFileName, number)
	err = ioutil.WriteFile(fileName, data, 0644)
	log.Println(fileName + " created!")
	return err
}

func searchAllComics(term string) {
	var wg sync.WaitGroup
	for i := 1; i <= comicNum; i++ {
		wg.Add(1)
		go func(number int, term string) {
			defer wg.Done()
			filepath := fmt.Sprintf(indexFileDir+"/"+indexFileName, number)
			// if file exist
			if _, err := os.Stat(filepath); err == nil {
				if err := searchComic(number, term); err != nil {
					log.Println(err)
				}
			}
		}(i, term)
	}
	wg.Wait()
}

func searchComic(number int, term string) error {
	c, err := loadComic(number)
	if err != nil {
		return err
	}
	if strings.Contains(c.SafeTitle, term) || strings.Contains(c.Transcript, term) {
		fmt.Printf(`<URL of Image>
%s
<Transcript>
%s

`, c.Img, c.Transcript)

	}
	return nil
}

func loadComic(number int) (*comic, error) {
	fileName := fmt.Sprintf(indexFileDir+"/"+indexFileName, number)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var c comic
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
