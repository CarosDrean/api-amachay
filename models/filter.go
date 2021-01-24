package models

type Filter struct {
	ID       string `json:"_id"`
	AuxID    string `json:"AuxId"`
	Type     string `json:"type"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}
