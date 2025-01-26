package database

import (
	"errors"
	"fmt"
	"interview/config"
	"interview/entity"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBProvider string

var (
	DBMySQL DBProvider = "mysql"
	DBTest  DBProvider = "test"
)

var (
	once sync.Once
	db   *gorm.DB

	ErrUnableToConnectToDatabase = errors.New("unable to connect to database")
	ErrDatabaseMigrationFailed   = errors.New("database migration failed")
	ErrDatabaseNil               = errors.New("database is nil")
	ErrInvalidDatabaseProvider   = errors.New("invalid database provider")
)

func InitDB(dbp DBProvider) error {
	var err error
	once.Do(func() {
		switch dbp {
		case DBMySQL:
			db, err = newMySQLDatabase()
		case DBTest:
			db, err = newTestDatabase()
		default:
			err = ErrInvalidDatabaseProvider
		}
	})

	return err
}

func Get() *gorm.DB {
	if db == nil {
		panic(ErrDatabaseNil)
	}

	return db
}

func newMySQLDatabase() (*gorm.DB, error) {
	cfg := config.Get().Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Join(ErrUnableToConnectToDatabase, err)
	}

	return migrate(gdb)
}

func newTestDatabase() (*gorm.DB, error) {
	gdb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, errors.Join(ErrUnableToConnectToDatabase, err)
	}

	return migrate(gdb)
}

func migrate(gdb *gorm.DB) (*gorm.DB, error) {
	if err := gdb.AutoMigrate(&entity.CartEntity{}, &entity.CartItem{}); err != nil {
		return nil, errors.Join(ErrDatabaseMigrationFailed, err)
	}
	return gdb, nil
}
