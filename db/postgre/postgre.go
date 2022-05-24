package postgre

import (
	"github.com/estiam/ticketing/db/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB = sqlite.DB

func New(dsn string) *DB {
	dsn = "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = sqlite.Migrate(conn)
	if err != nil {
		panic(err)
	}
	var db DB
	db.SetDB(conn)
	return &db
}
