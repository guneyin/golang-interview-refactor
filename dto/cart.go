package dto

import "interview/entity"

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}

type CartResponse struct {
	Error     string     `json:"Error"`
	CartItems []CartItem `json:"CartItems"`
}

type CartItem struct {
	ID       uint    `json:"ID"`
	Product  string  `json:"Product"`
	Quantity uint    `json:"Quantity"`
	Price    float64 `json:"Price"`
}

func (r *CartResponse) FromEntity(e entity.CartItems) {
	r.CartItems = make([]CartItem, len(e))
	for i, item := range e {
		r.CartItems[i] = CartItem{
			ID:       item.ID,
			Product:  item.ProductName,
			Quantity: item.Quantity,
			Price:    item.Price,
		}
	}
}
