package cart

import (
	"errors"
	"interview/database"
	"interview/entity"

	"gorm.io/gorm"
)

var (
	ErrInvalidItemName = errors.New("invalid item name")
	ErrCartNotFound    = errors.New("cart not found")
)

var productPriceList = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetCart(sessionID string) (entity.CartItems, error) {
	var cartItems entity.CartItems

	db := database.Get()

	cartID, err := r.initCart(sessionID, false)
	if err != nil {
		return nil, err
	}

	tx := db.Where("cart_id = ?", cartID).Find(&cartItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return cartItems, nil
}

func (r *Repository) AddItem(sessionID, product string, qty uint) error {
	cartID, err := r.initCart(sessionID, true)
	if err != nil {
		return err
	}

	price, err := getProductPrice(product)
	if err != nil {
		return err
	}

	cartItem := &entity.CartItem{}

	db := database.Get()
	tx := db.Where("cart_id = ? AND product_name = ?", cartID, product).First(&cartItem)

	if tx.Error != nil {
		switch {
		case errors.Is(tx.Error, gorm.ErrRecordNotFound):
		default:
			return tx.Error
		}
	}

	cartItem.CartID = cartID
	cartItem.ProductName = product
	cartItem.Quantity += qty
	cartItem.Price += float64(qty) * price

	return db.Save(cartItem).Error
}

func (r *Repository) DeleteItem(sessionID string, itemID uint) error {
	db := database.Get()

	cartID, err := r.initCart(sessionID, false)
	if err != nil {
		return err
	}

	return db.Where("cart_id = ? AND id = ?", cartID, itemID).Delete(&entity.CartItem{}).Error
}

func (r *Repository) initCart(sessionID string, create bool) (uint, error) {
	cart := &entity.CartEntity{}

	db := database.Get()
	tx := db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cart)
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.Error != nil {
		return 0, tx.Error
	}

	if create {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			cart.Status = entity.CartOpen
			cart.SessionID = sessionID

			tx = db.Save(&cart)
			if tx.Error != nil {
				return 0, tx.Error
			}
		}
	}

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return 0, ErrCartNotFound
	}

	return cart.ID, tx.Error
}

func getProductPrice(product string) (float64, error) {
	price, ok := productPriceList[product]
	if !ok {
		return 0, ErrInvalidItemName
	}

	return price, nil
}
