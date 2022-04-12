package main

import (
	"time"

	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/routes"
	"github.com/kimbu-chat/web-socket-manager-go/internal/pkg/server"
)

func main() {
	config.Init()
	defer func() {
		config.Close()
	}()

	router := routes.InitServer()
	//TODO: move address and timeout to config
	server.Run(":8080", router, 20*time.Second)
}
