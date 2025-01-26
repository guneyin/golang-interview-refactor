package cart

import (
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

func TestService(t *testing.T) {
	service := NewService()

	err := service.Add(sessionID, product, productQty)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Delete(sessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cart, err := service.GetCart(sessionID)
	if err != nil {
		t.Fatal(err)
	}
	if len(cart) > 0 {
		t.Fatal("cart should be nil")
	}
}
