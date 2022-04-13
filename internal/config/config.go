package config

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
)

func Init() {
	db.Connection()
	initGRPCCleint()
	initValidations()
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
	if err := closeGRPCConnection(); err != nil {
		panic(err)
	}
}
