package models

type Product struct {
	ID			int 	`json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	IdCategory  int     `json:"id_category"`
}

