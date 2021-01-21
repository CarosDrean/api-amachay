package models

type Provider struct {
	ID         int    `json:"_id"`
	IdBusiness int64  `json:"idBusiness"`
	Type       string `json:"type"`
}

type ProviderBusiness struct {
	ID         int    `json:"_id"`
	IdBusiness int    `json:"idBusiness"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	RUC        string `json:"ruc"`
	Address    string `json:"address"`
	Cel        string `json:"cel"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
}
