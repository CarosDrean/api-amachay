package models

type Product struct {
	ID          int     `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	IdCategory  int     `json:"idCategory"`
	Category    string  `json:"category"`
	Perishable  bool    `json:"perishable"`
	Ignore      bool    `json:"ignore"`
}

type ProductFill struct {
	ID               int     `json:"_id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Price            float64 `json:"price"`
	Stock            float64 `json:"stock"`
	IdProductMeasure int     `json:"idProductMeasure"`
	IdCategory       int     `json:"idCategory"`
	Category         string  `json:"category"`
	IdMeasure        int     `json:"idMeasure"`
	Measure          string  `json:"measure"`
	Unity            int     `json:"unity"`
	MinAlert         int     `json:"minAlert"`
	Perishable       bool    `json:"perishable"`
	Ignore           bool    `json:"ignore"`
}

