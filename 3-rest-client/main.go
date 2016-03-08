package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var r response

	resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("status code is %s\n", resp.Status)

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	json.Unmarshal(contents, &r)

	fmt.Printf("The movie : %s was released in %d - the IMDB rating is %d%% with %s votes\n",
		r.Title, r.Year, int(r.ImdbRating*10), r.ImdbVotes)
}
