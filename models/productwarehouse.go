package models

type ProductWarehouse struct {
	ID          int  `json:"_id"`
	IdProduct   int  `json:"idProduct"`
	IdWarehouse int  `json:"idWarehouse"`
	Ignore      bool `json:"ignore"`
}
