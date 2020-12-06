package models

type SystemUser struct {
	ID       int    `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IdPerson int64  `json:"idPerson"`
}

type UserResult struct {
	ID   string `json:"_id"`
	Role string `json:"role"`
}

type UserLogin struct {
	User     string
	Password string
}
