package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func init() {
	dsn := os.ExpandEnv("host=%DB_HOST user=%DB_USER password=%DB_PASSWORD dbname=golightweb port=5432")
	var err error
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Connection() *gorm.DB {
	return connection
}
