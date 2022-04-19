package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              Cfg.SentryDNS(),
		Release:          "0.1",
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("initSentry: %s", err)
	}

}

func flushSentry() {
	sentry.Flush(Cfg.SentryFlushTimeout())
}
