package models

type Product struct {
	ID          int     `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	IdCategory  int     `json:"idCategory"`
}
