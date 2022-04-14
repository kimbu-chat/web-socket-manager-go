package main

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/routes"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/server"
)

func main() {
	config.Init()
	defer func() {
		config.Close()
	}()

	app := routes.InitApp()
	//TODO: move address to config
	// https://github.com/kimbu-chat/web-socket-manager-go/issues/27
	server.Run(":8080", app)
}
