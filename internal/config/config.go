package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config/db"
)

var Cfg cfg

type cfg struct {
}

const DefaultSentryFlushTimeout = 5

func (c *cfg) SentryDNS() string {
	return os.Getenv("SENTRY_DNS")
}

func (c *cfg) SentryFlushTimeout() time.Duration {
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

func (c *cfg) DbHost() string {
	return os.Getenv("DB_HOST")
}

func (c *cfg) DbUser() string {
	return os.Getenv("DB_USER")
}

func (c *cfg) DbPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (c *cfg) DbName() string {
	return os.Getenv("DB_NAME")
}

func (c *cfg) DbPort() string {
	return os.Getenv("DB_PORT")
}

func Init() {
	loadCfg()
	initSentry()
	db.InitConnection(db.DbCfg{
		Host:     Cfg.DbHost(),
		User:     Cfg.DbUser(),
		Password: Cfg.DbPassword(),
		Name:     Cfg.DbName(),
		Port:     Cfg.DbPort(),
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

	Cfg = cfg{}
}
