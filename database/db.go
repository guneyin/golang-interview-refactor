package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"interview/entity"
	"sync"
)

type Database struct {
	*gorm.DB
}

var (
	dbOnce sync.Once
	db     *Database
)

func Get() *Database {
	dbOnce.Do(func() {
		var err error
		db, err = newDatabase()
		if err != nil {
			panic(err)
		}

	})
	return db
}

func newDatabase() (*Database, error) {
	dsn := "ice_user:9xz3jrd8wf@tcp(localhost:4001)/ice_db?charset=utf8mb4&parseTime=True&loc=Local"

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database := &Database{gdb}
	database.migrate()
	return database, nil
}

func (d *Database) migrate() {
	err := d.AutoMigrate(&entity.CartEntity{}, &entity.CartItem{})
	if err != nil {
		panic(err)
	}
}
