package models

type User struct {
	ID          int    `json:"_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"` // Admin or User
	IdPerson    int64  `json:"idPerson"`
	IdWarehouse int    `json:"idWarehouse"`
}

type ClaimUser struct {
	ID   string `json:"_id"`
	Role string `json:"role"`
}

type Login struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
