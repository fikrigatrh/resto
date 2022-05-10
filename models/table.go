package models

type Table struct {
	Meta
	TableNumber int    `json:"tableNumber"`
	Status      string `json:"status"`
}
