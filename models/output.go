package models

type Output struct {
	ID        int    `json:"_id"`
	Date      string `json:"date"`
	Quantity  int    `json:"quantity"`
	IdProduct int    `json:"idProduct"`
}
