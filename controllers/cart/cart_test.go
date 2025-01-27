package cart_test

import (
	"github.com/gin-gonic/gin"
	"interview/controllers/cart"
	"interview/mw"
	"interview/test"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var (
	router  *gin.Engine
	handler *cart.Handler
)

func initTest() {
	test.ChangeWorkDir()
	test.InitTestDB()

	router = test.NewTestRouter()
	router.Use(mw.UseSession())

	handler = cart.NewHandler()
}

func TestIndex(t *testing.T) {
	initTest()

	t.Run("Get Index Page", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		router.GET("/", handler.Index)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(recorder, ctx.Request)
		if recorder.Code != http.StatusOK {
			t.Fatalf("Want %d, got %d", http.StatusOK, recorder.Code)
		}
	})

}

func TestAddItem(t *testing.T) {
	initTest()

	t.Run("Add Item", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		form := url.Values{}
		form.Add("product", test.Product)
		form.Add("quantity", strconv.Itoa(test.ProductQty))

		router.POST("/add-item", handler.AddItem)

		ctx.Request = httptest.NewRequest(http.MethodPost, "/add-item", strings.NewReader(form.Encode()))
		ctx.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		router.ServeHTTP(recorder, ctx.Request)
		if recorder.Code != http.StatusFound {
			t.Fatalf("Want %d, got %d", http.StatusFound, recorder.Code)
		}
	})

	t.Run("Get Cart With Item", TestIndex)
}

func TestDeleteItem(t *testing.T) {
	initTest()

	t.Run("Add Item", TestAddItem)

	t.Run("Delete Item", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		router.GET("/remove-cart-item", handler.RemoveCartItem)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/remove-cart-item?cart_item_id=1", nil)

		router.ServeHTTP(recorder, ctx.Request)
		if recorder.Code != http.StatusFound {
			t.Fatalf("Want %d, got %d", http.StatusFound, recorder.Code)
		}
	})
}
