package main

import (
	"github.com/kimbu-chat/web-socket-manager-go/internal/config"
	"github.com/kimbu-chat/web-socket-manager-go/internal/config/routes"
)

func main() {
	config.InitGRPCCleint()
	defer func() {
		if err := config.CloseGRPCConnection(); err != nil {
			panic(err)
		}
	}()

	server := routes.InitServer()

	if err := server.Run(); err != nil {
		panic(err)
	}
}
