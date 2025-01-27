package cart_test

import (
	"interview/database"
	"interview/services/cart"
	"testing"

	"github.com/google/uuid"
)

const (
	product        = "shoe"
	invalidProduct = "t-shirt"
	productQty     = 2
)

var sessionID = uuid.New().String()

func initDB() {
	err := database.InitDB(database.DBTest)
	if err != nil {
		panic(err)
	}
}

func TestService(t *testing.T) {
	initDB()
	service := cart.NewService()

	err := service.Add(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Delete(sessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err := service.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}
	if len(cartItems) > 0 {
		t.Fatal("cart should be nil")
	}

	err = service.Add(sessionID, invalidProduct, productQty)
	if err == nil {
		t.Fatal("cart should not be created")
	}
}
