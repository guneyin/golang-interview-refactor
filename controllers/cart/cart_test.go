package cart_test

import (
	"interview/controllers/cart"
	"interview/mw"
	"interview/test"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	router  *gin.Engine
	handler *cart.Handler

	paramProduct    string
	paramProductQty int
	paramStatusCode int
)

func setParams(product string, qty, code int) {
	paramProduct = product
	paramProductQty = qty
	paramStatusCode = code
}

func initTest() {
	test.ChangeWorkDir()
	test.InitTestDB()

	router = test.NewTestRouter()
	router.Use(mw.UseSession())

	handler = cart.NewHandler()
}

func TestIndex(t *testing.T) {
	initTest()

	t.Run("Get Index Page", index)
}

func TestAddItem(t *testing.T) {
	initTest()

	setParams(test.Product, test.ProductQty, http.StatusFound)
	t.Run("Add Item", addItem)

	t.Run("Get Cart With Item", TestIndex)
}

func TestAddItemInvalidQty(t *testing.T) {
	initTest()

	setParams(test.Product, 0, http.StatusSeeOther)
	t.Run("Add Item With Invalid Qty", addItem)
}

func TestAddItemInvalidProduct(t *testing.T) {
	initTest()

	setParams(test.InvalidProduct, test.ProductQty, http.StatusSeeOther)
	t.Run("Add Item With Invalid Product", addItem)
}

func TestDeleteItem(t *testing.T) {
	initTest()

	setParams(test.Product, test.ProductQty, http.StatusFound)
	t.Run("Add Item", addItem)

	t.Run("Delete Item", deleteItem)
}

func index(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	router.GET("/", handler.Index)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(recorder, ctx.Request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("Want %d, got %d", http.StatusOK, recorder.Code)
	}
}

func addItem(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	form := url.Values{}
	form.Add("product", paramProduct)
	form.Add("quantity", strconv.Itoa(paramProductQty))

	router.POST("/add-item", handler.AddItem)

	ctx.Request = httptest.NewRequest(http.MethodPost, "/add-item", strings.NewReader(form.Encode()))
	ctx.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(recorder, ctx.Request)
	if recorder.Code != paramStatusCode {
		t.Fatalf("Want %d, got %d", paramStatusCode, recorder.Code)
	}
}

func deleteItem(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	router.GET("/remove-cart-item", handler.RemoveCartItem)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/remove-cart-item?cart_item_id=1", nil)

	router.ServeHTTP(recorder, ctx.Request)
	if recorder.Code != http.StatusFound {
		t.Fatalf("Want %d, got %d", http.StatusFound, recorder.Code)
	}
}
