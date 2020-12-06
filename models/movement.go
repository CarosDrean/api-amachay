package models

type Movement struct {
	ID          int    `json:"_id"`
	IdProduct   int    `json:"idPerson"`
	IdWarehouse int    `json:"idWarehouse"`
	IdUser      int    `json:"idUser"`
	IdClient    int    `json:"idClient"`
	Date        string `json:"date"`
	Quantity    int    `json:"quantity"`
	Type        string `json:"type"`
}
