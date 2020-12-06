package models

type Warehouse struct {
	ID      int    `json:"_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	State   bool   `json:"state"`
}
