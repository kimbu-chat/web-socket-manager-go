package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.ExpandEnv("host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME port=$DB_PORT")
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
}

func Connection() *gorm.DB {
	return connection
}

func SQLDB() (*sql.DB, error) {
	return Connection().DB()
}

func Close() error {
	sqlDb, err := Connection().DB()
	if err != nil {
		return err
	}

	return sqlDb.Close()
}
