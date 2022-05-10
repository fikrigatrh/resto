package models

type Menu struct {
	Meta
	Name  string `json:"name"`
	Price int    `json:"price"`
}
