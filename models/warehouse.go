package models

type Warehouse struct {
	ID      int    `json:"_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	State   string `json:"state"`
}
