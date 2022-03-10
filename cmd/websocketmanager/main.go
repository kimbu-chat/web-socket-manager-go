package main

import "github.com/kimbu-chat/web-socket-manager-go/internal/config"

func main() {
	server := config.InitServer()
	server.Run()
}
