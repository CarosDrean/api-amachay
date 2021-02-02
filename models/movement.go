package models

type Movement struct {
	ID          int     `json:"_id"`
	IdProduct   int     `json:"idProduct"`
	Product     string  `json:"product"`
	Measure     string  `json:"measure"`
	IdWarehouse int     `json:"idWarehouse"`
	IdUser      int     `json:"idUser"`
	IdClient    int     `json:"idClient"`
	IdProvider  int     `json:"idProvider"`
	Date        string  `json:"date"`
	Quantity    float32 `json:"quantity"`
	Type        string  `json:"type"`

	IdLot      int    `json:"idLot"`
	IdBrand    int    `json:"idBrand"`
	Lot        string `json:"lot"`
	Brand      string `json:"brand"`
	DueDate    string `json:"dueDate"`
	Perishable bool   `json:"perishable"`
}
