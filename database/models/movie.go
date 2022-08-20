package model

type Movie struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Genre       string `json: "genre"`
	Description string `json: "description"`
	Image       string `json: "image_url"`
	Source      string `json: "source"`
	Language    string `json: "language"`
	Duration    string `json: "duration"`
	Category    string `json: "category"`
}
