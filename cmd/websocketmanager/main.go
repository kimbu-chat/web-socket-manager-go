package main

import "github.com/macmaczhl/websocketmanager/internal/config"

func main() {
	server := config.InitServer()
	server.Run()
}
