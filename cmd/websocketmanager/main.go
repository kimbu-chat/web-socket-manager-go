package main

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/routes"
)

func main() {
	config.InitGRPCCleint()
	defer config.CloseGRPCConnection()

	server := routes.InitServer()
	server.Run()
}
