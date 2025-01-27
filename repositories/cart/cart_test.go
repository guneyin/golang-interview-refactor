package cart_test

import (
	"interview/repositories/cart"
	"interview/test"
	"testing"
)

func TestRepository_GetCart(t *testing.T) {
	test.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Get Empty Cart", func(t *testing.T) {
		_, err := repo.GetCart(test.SessionID)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Get Cart With Item", func(t *testing.T) {
		err := repo.AddItem(test.SessionID, test.Product, test.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err := repo.GetCart(test.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 1 {
			t.Fatal("item count should be 1")
		}

		productPrice, err := cart.GetProductPrice(test.Product)
		if err != nil {
			t.Fatal(err)
		}

		itemPrice := productPrice * test.ProductQty

		item := cartItems[0]
		if item.Price != itemPrice {
			t.Fatal("price should be equal")
		}
	})
}

func TestRepository_AddItem(t *testing.T) {
	test.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Add Item", func(t *testing.T) {
		err := repo.AddItem(test.SessionID, test.Product, test.ProductQty)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Add Item With Invalid Product", func(t *testing.T) {
		err := repo.AddItem(test.SessionID, test.InvalidProduct, test.ProductQty)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestRepository_DeleteItem(t *testing.T) {
	test.InitTestDB()
	repo := cart.NewRepository()

	t.Run("Delete Item", func(t *testing.T) {
		err := repo.AddItem(test.SessionID, test.Product, test.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err := repo.GetCart(test.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 1 {
			t.Fatal("item count should be 1")
		}

		err = repo.DeleteItem(test.SessionID, 1)
		if err != nil {
			t.Fatal(err)
		}

		cartItems, err = repo.GetCart(test.SessionID)
		if err != nil {
			t.Fatal(err)
		}

		if len(cartItems) != 0 {
			t.Fatal("item count should be 0")
		}
	})

	t.Run("Delete Invalid Session Item", func(t *testing.T) {
		err := repo.AddItem(test.SessionID, test.Product, test.ProductQty)
		if err != nil {
			t.Fatal(err)
		}

		err = repo.DeleteItem(test.InvalidSessionID, 1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(err)
	})
}
