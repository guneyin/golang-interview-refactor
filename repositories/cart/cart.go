package cart

import (
	"errors"
	"gorm.io/gorm"
	"interview/database"
	"interview/entity"
)

var ErrInvalidItemName = errors.New("invalid item name")

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
	var (
		cartEntity entity.CartEntity
		cartItems  entity.CartItems
	)

	db := database.Get()
	tx := db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cartEntity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = db.Where("cart_id = ?", cartEntity.ID).Find(&cartItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return cartItems, nil
}

func (r *Repository) AddItem(sessionID, product string, qty uint) error {
	cartID, err := r.initCart(sessionID)
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
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.Error != nil {
		return err
	}

	cartItem.CartID = cartID
	cartItem.ProductName = product
	cartItem.Quantity += qty
	cartItem.Price += float64(qty) * price

	return db.Save(cartItem).Error
}

func (r *Repository) DeleteItem(sessionID string, itemID uint) error {
	cart := &entity.CartEntity{}

	db := database.Get()
	tx := db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cart)
	if tx.Error != nil {
		return tx.Error
	}

	cartItem := &entity.CartItem{
		Model: gorm.Model{ID: itemID},
	}

	tx = db.First(cartItem)
	if tx.Error != nil {
		return tx.Error
	}

	return db.Delete(&cartItem).Error
}

func (r *Repository) initCart(sessionID string) (uint, error) {
	cartEntity := &entity.CartEntity{}

	db := database.Get()
	tx := db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cartEntity)
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.Error != nil {
		return 0, tx.Error
	}

	cartEntity.Status = entity.CartOpen
	cartEntity.SessionID = sessionID

	tx = db.Save(&cartEntity)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return cartEntity.ID, nil
}

func getProductPrice(product string) (float64, error) {
	price, ok := productPriceList[product]
	if !ok {
		return 0, ErrInvalidItemName
	}

	return price, nil
}
