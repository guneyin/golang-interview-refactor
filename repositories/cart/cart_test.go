package cart_test

import (
	"errors"
	"interview/database"
	"interview/repositories/cart"
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

func TestRepository_GetCart(t *testing.T) {
	initDB()

	repo := cart.NewRepository()

	_, err := repo.GetCart(sessionID)
	if !errors.Is(err, cart.ErrCartNotFound) {
		t.Fatal("expected ErrCartNotFound")
	}

	err = repo.AddItem(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err := repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cartItems) != 1 {
		t.Fatal("item count should be 1")
	}

	productPrice, err := cart.GetProductPrice(product)
	if err != nil {
		t.Fatal(err)
	}

	itemPrice := productPrice * productQty

	item := cartItems[0]
	if item.Price != itemPrice {
		t.Fatal("price should be equal")
	}
}

func TestRepository_AddItem(t *testing.T) {
	initDB()

	repo := cart.NewRepository()

	err := repo.AddItem(sessionID, invalidProduct, productQty)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRepository_DeleteItem(t *testing.T) {
	initDB()

	repo := cart.NewRepository()

	err := repo.AddItem(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err := repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cartItems) != 1 {
		t.Fatal("item count should be 1")
	}

	err = repo.DeleteItem(sessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err = repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cartItems) != 0 {
		t.Fatal("item count should be 0")
	}
}
