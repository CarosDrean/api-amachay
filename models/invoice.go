package models

type Invoice struct {
	ID      int    `json:"_id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Date    string `json:"date"`
	IdImage string `json:"idImage"`
}
