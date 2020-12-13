package models

type UserPerson struct {
	ID          int    `json:"_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	IdPerson    int64  `json:"idPerson"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Cel         string `json:"cel"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Dni         string `json:"dni"`
	Mail        string `json:"mail"`
	IdWarehouse int    `json:"idWarehouse"`
}

type ClientPerson struct {
	ID       int    `json:"_id"`
	Type     string `json:"type"`
	IdPerson int64  `json:"idPerson"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Cel      string `json:"cel"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Dni      string `json:"dni"`
	Mail     string `json:"mail"`
}
