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

type Repository interface {
	GetCart(sessionID string) (entity.CartItems, error)
	AddItem(sessionID, productID string, qty uint) error
	DeleteItem(sessionID string, itemID uint) error
}

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetCart(sessionID string) (entity.CartItems, error) {
	db := database.Get()
	var (
		cartEntity entity.CartEntity
		cartItems  entity.CartItems
	)

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

func (r *repositoryImpl) AddItem(sessionID, product string, qty uint) error {
	cartID, err := r.initCart(sessionID)
	if err != nil {
		return err
	}

	price, err := getProductPrice(product)
	if err != nil {
		return err
	}

	db := database.Get()

	found := &entity.CartItem{}
	_ = db.Where("cart_id = ? AND product_name = ?", cartID, product).First(&found)
	found.CartID = cartID
	found.ProductName = product
	found.Quantity += qty
	found.Price += float64(qty) * price

	return db.Save(found).Error
}

func (r *repositoryImpl) DeleteItem(sessionID string, itemID uint) error {
	db := database.Get()

	cart := &entity.CartEntity{}
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

func (r *repositoryImpl) initCart(sessionID string) (uint, error) {
	db := database.Get()

	cartEntity := &entity.CartEntity{}
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
