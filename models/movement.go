package models

type Movement struct {
	ID          int     `json:"_id"`
	IdProduct   int     `json:"idProduct"`
	Product     string  `json:"product"`
	IdWarehouse int     `json:"idWarehouse"`
	IdUser      int     `json:"idUser"`
	IdClient    int     `json:"idClient"`
	Date        string  `json:"date"`
	Quantity    float32 `json:"quantity"`
	Type        string  `json:"type"`
}
