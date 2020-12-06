package models

type Person struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Cel      string `json:"cel"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Dni      string `json:"dni"`
	Mail     string `json:"mail"`
}
