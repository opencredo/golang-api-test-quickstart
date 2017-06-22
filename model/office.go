package model

type Office struct {
	Dept     string `json:"dept"`
	Position string `json:"position"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
	Person   string `json:""`
	Source   string `json:""`
}