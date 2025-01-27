package test

import (
	"github.com/google/uuid"
	"interview/config"
	"interview/database"
	"interview/router"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	Product        = "shoe"
	InvalidProduct = "t-shirt"
	ProductQty     = 2
)

var (
	SessionID        = uuid.New().String()
	InvalidSessionID = uuid.New().String()
)

func InitTestDB() {
	err := database.InitDB(database.DBTest)
	if err != nil {
		panic(err)
	}
}

func ChangeWorkDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func NewTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	cfg := config.AppConfig{
		SessionSecret: "s€cR€t",
		Port:          "8088",
	}

	return router.NewRouter(cfg)
}
