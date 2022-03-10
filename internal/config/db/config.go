package db

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func init() {
	godotenv.Load()

	dsn := os.ExpandEnv("host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=websocketmanager port=5439")
	var err error
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Connection() *gorm.DB {
	return connection
}
