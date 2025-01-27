package cart_test

import (
	"interview/services/cart"
	"interview/test"
	"testing"
)

func TestService(t *testing.T) {
	test.InitTestDB()

	service := cart.NewService()

	err := service.Add(test.SessionID, test.Product, test.ProductQty)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Delete(test.SessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err := service.GetCart(test.SessionID)
	if err != nil {
		t.Fatal(err)
	}
	if len(cartItems) > 0 {
		t.Fatal("cart should be nil")
	}

	err = service.Add(test.SessionID, test.InvalidProduct, test.ProductQty)
	if err == nil {
		t.Fatal("cart should not be created")
	}
}
