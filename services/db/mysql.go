package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	MySQL *gorm.DB
)

func ConnectMySQL(databaseType, dsn string) {
	log.Printf("Connecting to %s database on %s\n", databaseType, dsn)
	database, err := gorm.Open(databaseType, dsn)
	if err != nil {
		panic(err)
	}
	MySQL = database
}

func AutoMigration(model interface{}) {
	log.Println("Migration database model")
	MySQL.AutoMigrate(model)
}
