package entity

import (
	"gorm.io/gorm"
)

const (
	CartOpen   = "open"
	CartClosed = "closed"
)

type CartEntity struct {
	gorm.Model
	Total     float64
	SessionID string
	Status    string
}

type CartItem struct {
	gorm.Model
	CartID      uint
	ProductName string
	Quantity    int
	Price       float64
}
