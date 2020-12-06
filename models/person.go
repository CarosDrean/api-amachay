package models

type Person struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    int    `json:"phone"`
	Adress  string `json:"adress"`
	Dni      int    `json:"dni"`
	Mail     string `json:"mail"`
}
