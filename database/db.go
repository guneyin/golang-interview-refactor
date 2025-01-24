package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"interview/entity"
)

func GetDatabase() *gorm.DB {
	// MySQL connection string
	// Update the username, password, host, port, and database name accordingly
	dsn := "ice_user:9xz3jrd8wf@tcp(localhost:4001)/ice_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate() {
	db := GetDatabase()

	// AutoMigrate will create or update the tables based on the models
	err := db.AutoMigrate(&entity.CartEntity{}, &entity.CartItem{})
	if err != nil {
		panic(err)
	}
}
