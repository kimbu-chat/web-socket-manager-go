package main

import (
	"github.com/getsentry/sentry-go"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/routes"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/server"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			hub := sentry.CurrentHub().Clone()
			hub.Recover(err)
			hub.Flush(config.SentryFlushTimeout())
			panic(err)
		}
	}()

	config.Init()
	defer config.Close()

	app := routes.InitApp()

	server.Run(config.ListeningAddress(), app)
}
