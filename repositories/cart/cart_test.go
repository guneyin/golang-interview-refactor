package cart_test

import (
	"interview/repositories/cart"
	"interview/testutils"
	"testing"
)

func TestRepository_GetCart(t *testing.T) {
	testutils.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Get Empty Cart", func(t *testing.T) {
		_, err := repo.GetCart(testutils.SessionID)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get Cart With Item", func(t *testing.T) {
		err := repo.AddItem(testutils.SessionID, testutils.Product, testutils.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err := repo.GetCart(testutils.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 1 {
			t.Fatal("item count should be 1")
		}

		productPrice, err := cart.GetProductPrice(testutils.Product)
		if err != nil {
			t.Fatal(err)
		}

		itemPrice := productPrice * testutils.ProductQty

		item := cartItems[0]
		if item.Price != itemPrice {
			t.Fatal("price should be equal")
		}
	})
}

func TestRepository_AddItem(t *testing.T) {
	testutils.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Add Item", func(t *testing.T) {
		err := repo.AddItem(testutils.SessionID, testutils.Product, testutils.ProductQty)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Add Item With Invalid Product", func(t *testing.T) {
		err := repo.AddItem(testutils.SessionID, testutils.InvalidProduct, testutils.ProductQty)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestRepository_DeleteItem(t *testing.T) {
	testutils.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Delete Item", func(t *testing.T) {
		err := repo.AddItem(testutils.SessionID, testutils.Product, testutils.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err := repo.GetCart(testutils.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 1 {
			t.Fatal("item count should be 1")
		}

		err = repo.DeleteItem(testutils.SessionID, 1)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err = repo.GetCart(testutils.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 0 {
			t.Fatal("item count should be 0")
		}
	})

	t.Run("Delete Invalid Session Item", func(t *testing.T) {
		err := repo.AddItem(testutils.SessionID, testutils.Product, testutils.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		err = repo.DeleteItem(testutils.InvalidSessionID, 1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(err)
	})
}
