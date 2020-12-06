package models

type SystemUser struct {
	ID       int    `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	IdPerson int64  `json:"personId"`
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
	PersonID int64  `json:"personId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Phone    int    `json:"phone"`
	Adress   string `json:"adress"`
	Dni      int    `json:"dni"`
	Mail     string `json:"mail"`
}
