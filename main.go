package main

import (
	"client-server/config"
	"client-server/src"
)

var conf config.Config = config.LoadConfig(src.ConfigPath)

func main() {
	var server *src.Server
	src.CreateServer(&server)

	go src.StartListening(&server)

	if conf.ReadFromFile {
		server.HandleExample()
	} else {
		<-server.CloseChan
		server.CloseServer()
	}
}
