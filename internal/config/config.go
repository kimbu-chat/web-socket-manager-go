package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
)

const DefaultSentryFlushTimeout = 5

func SentryDNS() string {
	return os.Getenv("SENTRY_DNS")
}

func SentryFlushTimeout() time.Duration {
	tStr, ok := os.LookupEnv("SENTRY_FLUSH_TIMEOUT")
	if !ok {
		return DefaultSentryFlushTimeout * time.Second
	}
	seconds, err := strconv.Atoi(tStr)
	if err != nil {
		return DefaultSentryFlushTimeout * time.Second
	}

	return time.Duration(seconds) * time.Second
}

func DbHost() string {
	return os.Getenv("DB_HOST")
}

func DbUser() string {
	return os.Getenv("DB_USER")
}

func DbPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func DbName() string {
	return os.Getenv("DB_NAME")
}

func DbPort() string {
	return os.Getenv("DB_PORT")
}

func Init() {
	loadCfg()
	initSentry()
	db.InitConnection(db.DbCfg{
		Host:     DbHost(),
		User:     DbUser(),
		Password: DbPassword(),
		Name:     DbName(),
		Port:     DbPort(),
	})
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

	flushSentry()
}

func loadCfg() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
