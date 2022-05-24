package sqlite

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/estiam/ticketing/db"
	"github.com/estiam/ticketing/model"
)

type DB struct {
	conn *gorm.DB
}

var _ db.Store = &DB{}

// func New(dbName string) *DB {
// 	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = Migrate(conn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &DB{
// 		conn: conn,
// 	}
// }

func NewPostgre(dsn string) *DB {
	dsn = "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Migrate(conn)
	if err != nil {
		panic(err)
	}

	return &DB{
		conn: conn,
	}
}

func (db *DB) SetDB(conn *gorm.DB) {
	db.conn = conn
}

func Migrate(conn *gorm.DB) error {
	err := conn.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(&model.Ticket{})
	if err != nil {
		return err
	}

	return nil
}
