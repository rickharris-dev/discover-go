package main

type movieList struct {
	Search []struct {
		ImdbID string
	}
}

type movieData struct {
	Title      string
	Year       string
	ImdbRating float64 `json:"imdbRating,string"`
	ImdbVotes  string
}
