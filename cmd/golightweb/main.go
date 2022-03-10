package main

import "github.com/macmaczhl/golightweb/internal/config"

func main() {
	server := config.InitServer()
	server.Run()
}
