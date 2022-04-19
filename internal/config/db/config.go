package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

type DbCfg struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func InitConnection(cfg DbCfg) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	var err error
	connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
