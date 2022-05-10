package models

type TransDetails struct {
	Meta
	MenuName string `json:"menuName"`
	QTY      int    `json:"qty"`
	Price    int    `json:"price"`
	Total    int    `json:"total"`
}

type TransHeader struct {
	TransID     string         `json:"transID"`
	TableNumber string         `json:"tableNumber"`
	TotalPrice  int            `json:"totalPrice"`
	Details     []TransDetails `json:"details"`
}

type Transaction struct {
	Meta
	TableID string `json:"tableID"`
	Date    string `json:"date"`
}

type Order struct {
	MenuID string `json:"menuID"`
	QTY    int    `json:"qty"`
	Price  int    `json:"price"`
	Notes  string `json:"notes"`
}
