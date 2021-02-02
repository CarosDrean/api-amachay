package models

type Lot struct {
	ID      int    `json:"_id"`
	Lot     string `json:"lot"`
	Brand   string `json:"brand"`
	DueDate string `json:"dueDate"`
}
