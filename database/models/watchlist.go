package model

type Watchlist struct {
	MovieId string `json:"movie_id"`
	UserId  string `json:"user_id,omitempty"`
}
