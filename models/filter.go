package models

type Filter struct {
	ID       string `json:"_id"`
	Type     string `json:"type"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}
