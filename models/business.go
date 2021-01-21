package models

type Business struct {
	ID      int    `json:"_id"`
	Name    string `json:"name"`
	RUC     string `json:"ruc"`
	Address string `json:"address"`
	Cel     string `json:"cel"`
	Phone   string `json:"phone"`
	Mail    string `json:"mail"`
}
