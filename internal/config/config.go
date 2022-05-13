package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

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

func ListeningPort() string {
	return os.Getenv("LISTENING_PORT")
}

func ListeningHost() string {
	return os.Getenv("LISTENING_HOST")
}

func ListeningAddress() string {
	return fmt.Sprintf("%s:%s", ListeningHost(), ListeningPort())
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

func CentrifugoGRPCAddress() string {
	return os.Getenv("CENTRIFUGO_GRPC_ADDRESS")
}

func CentrifugoAPIKey() string {
	return os.Getenv("CENTRIFUGO_API_KEY")
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
		logrus.Info(".env file is not loaded. Environemnt variables are used")
	}
}
