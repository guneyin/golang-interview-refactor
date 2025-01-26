package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"interview/config"
	"interview/entity"
	"sync"
)

var (
	once sync.Once
	db   *Database
)

type Database struct {
	*gorm.DB
}

func Get() *Database {
	once.Do(func() {
		var err error
		db, err = newDatabase()
		if err != nil {
			panic(err)
		}

	})
	return db
}

func newDatabase() (*Database, error) {
	cfg := config.Get().Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

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
