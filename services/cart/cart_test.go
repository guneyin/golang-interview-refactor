package cart_test

import (
	"interview/services/cart"
	"interview/testutils"
	"testing"
)

func TestService(t *testing.T) {
	testutils.InitTestDB()

	service := cart.NewService()

	err := service.Add(testutils.SessionID, testutils.Product, testutils.ProductQty)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Delete(testutils.SessionID, 1)
	if err != nil {
		t.Fatal(err)
	}

	cartItems, err := service.GetCart(testutils.SessionID)
	if err != nil {
		t.Fatal(err)
	}
	if len(cartItems) > 0 {
		t.Fatal("cart should be nil")
	}

	err = service.Add(testutils.SessionID, testutils.InvalidProduct, testutils.ProductQty)
	if err == nil {
		t.Fatal("cart should not be created")
	}
}
