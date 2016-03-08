package main

type response struct {
	Title      string
	Year       int `json:"Year,string"`
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  int     `json:"Metascore,string"`
	ImdbRating float64 `json:"imdbRating,string"`
	ImdbVotes  string
	ImdbID     string
	Type       string
	Response   bool
}
