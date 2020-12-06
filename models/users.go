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

type UserPerson struct {
	ID       int    `json:"_id"`
	IdPerson int64  `json:"idPerson"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
	Dni      int    `json:"dni"`
	Mail     string `json:"mail"`
}
