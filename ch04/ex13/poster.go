package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type SearchResult struct {
	Movies []Movie `json:"Search"`
}

type Movie struct {
	Title  string
	ImdbID string
	Poster string
}

const apiURL = "http://www.omdbapi.com/"

func main() {
	token := os.Getenv("TOKEN")
	if token == "" || len(os.Args) < 2 {
		printUsage()
		return
	}
	term := os.Args[1]

	// search movie
	res, err := searchMovie(token, term)
	handleErr(err)

	// define target movie
	var target Movie
	if len(res.Movies) == 0 {
		fmt.Println("Movie not found!")
		return
	} else if len(res.Movies) == 1 {
		target = res.Movies[0]
	} else {
		target = handleMultiMovies(res.Movies)
	}

	// download poster
	err = savePoster(target)
	handleErr(err)
	fmt.Println("Download Poster of \"" + target.Title + "\" Completed!!")
}

func printUsage() {
	fmt.Println(`USAGE:
	TOKEN=xxxx go run poster.go [SEARCH TERM]

TOKEN:
	please set to ENV variable your Open Movie Database API Key`)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func searchMovie(token, term string) (*SearchResult, error) {
	searchMovieURL := apiURL + "?apikey=" + token + "&s=" + url.QueryEscape(term)
	resp, err := http.Get(searchMovieURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search movie failed: %s\n%s", resp.Status, searchMovieURL)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func handleMultiMovies(movies []Movie) Movie {
	fmt.Println("ID: Title")
	fmt.Println("----------------------------------------")
	for i, m := range movies {
		fmt.Printf("%2d: %s\n", i, m.Title)
	}

	var id int
	fmt.Println("multi movies found. Choose ID to Download Poster.")
	fmt.Print("> ")
	for {
		_, err := fmt.Scan(&id)
		if err == nil && id < len(movies) {
			break
		} else {
			fmt.Printf("Invalid number. Please input 0 to %d\n", len(movies))
		}
	}
	return movies[id]
}

func savePoster(target Movie) error {
	posterURL := target.Poster
	fileName := strings.ReplaceAll(target.Title, " ", "_") + filepath.Ext(target.Poster)
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch poster failed: %s\n%s", resp.Status, posterURL)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, data, 0644)
	return err
}
