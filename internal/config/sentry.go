package config

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              SentryDNS(),
		Release:          "0.1",
		AttachStacktrace: true,
	})
	if err != nil {
		logrus.Fatalf("initSentry: %s", err)
	}

}

func flushSentry() {
	sentry.Flush(SentryFlushTimeout())
}
