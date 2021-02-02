package models

type Lot struct {
	ID      int    `json:"_id"`
	Lot     string `json:"lot"`
	DueDate string `json:"dueDate"`
}
