package models

type Client struct {
	ID       int    `json:"_id"`
	IdPerson int64  `json:"idPerson"`
	Type     string `json:"type"`
}
