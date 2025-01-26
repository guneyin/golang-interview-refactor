package cart

import (
	"errors"
	"github.com/google/uuid"
	"interview/database"
	"testing"
)

const (
	product        = "shoe"
	invalidProduct = "t-shirt"
	productQty     = 2
)

var sessionID = uuid.New().String()

func init() {
	err := database.InitDB(database.DBTest)
	if err != nil {
		panic(err)
	}
}

func TestRepository_GetCart(t *testing.T) {
	repo := NewRepository()

	_, err := repo.GetCart(sessionID)
	if !errors.Is(err, ErrCartNotFound) {
		t.Fatal("expected ErrCartNotFound")
	}

	err = repo.AddItem(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	cart, err := repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cart) != 1 {
		t.Fatal("item count should be 1")
	}

	productPrice, err := getProductPrice(product)
	if err != nil {
		t.Fatal(err)
	}

	itemPrice := productPrice * productQty

	item := cart[0]
	if item.Price != itemPrice {
		t.Fatal("price should be equal")
	}
}

func TestRepository_AddItem(t *testing.T) {
	repo := NewRepository()

	err := repo.AddItem(sessionID, invalidProduct, productQty)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRepository_DeleteItem(t *testing.T) {
	repo := NewRepository()

	err := repo.AddItem(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	cart, err := repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cart) != 1 {
		t.Fatal("item count should be 1")
	}

	err = repo.DeleteItem(sessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cart, err = repo.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}

	if len(cart) != 0 {
		t.Fatal("item count should be 0")
	}
}
