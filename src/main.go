package main

import "clean-architecture-go/infrastructure/waf/echo/server"

func main() {
	server.NewServer().Run()
}
