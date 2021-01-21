package models

type ProductMeasure struct {
	ID        int `json:"_id"`
	IdProduct int `json:"idProduct"`
	IdMeasure int `json:"idMeasure"`
	Unity     int `json:"unity"`
}
