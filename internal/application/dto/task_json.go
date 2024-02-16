package dto

type Task struct {
	ID         string `json:"id"`
	Difficulty int    `json:"difficulty"`
	Duration   int    `json:"duration"`
}
