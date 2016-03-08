package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	var movie string
	var l movieList
	var d movieData
	site := "http://www.omdbapi.com/?s="
	page := "&page="

	flag.StringVar(&movie, "movie", "Batman", "Enter name of movie to search")
	flag.Parse()

	url := []string{site, movie, page}
	resp, err := http.Get(strings.Join(url, ""))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(contents, &l)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	for i := 0; i < len(l.Search); i++ {
		urlM := []string{"http://www.omdbapi.com/?i=", l.Search[i].ImdbID}
		r, err := http.Get(strings.Join(urlM, ""))
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		err = json.Unmarshal(result, &d)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("The movie : %s was released in %s - the IMDB rating is %d%% with %s votes.\n",
			d.Title, d.Year, int(d.ImdbRating*10), d.ImdbVotes)
	}
	t2 := time.Now()
	fmt.Printf("execution time is %v\n", t2.Sub(t1))
}
