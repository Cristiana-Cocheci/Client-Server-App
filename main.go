package main

import (
	"client-server/config"
	"client-server/src"
	"fmt"
	"net"
	"sync"
)

var conf config.Config = config.LoadConfig(src.ConfigPath)

func handleExample(server *src.Server) {
	wg := sync.WaitGroup{}
	connList := make(map[int]net.Conn)

	wg.Add(src.ExampleClientsNumber)
	for i := 0; i < src.ExampleClientsNumber; i++ {
		go func() {
			fmt.Println("New connection.")
			defer wg.Done()
			connList[i] = server.ConnectToServer(i, conf.ReadFromFile)
		}()
	}
	wg.Wait()

	wg.Add(src.ExampleClientsNumber)
	for i := 0; i < src.ExampleClientsNumber; i++ {
		go func() {
			defer wg.Done()
			req_number := src.GetReqNumber(i + 1)
			for j := 0; j < req_number; j++ {
				server.SendRequestToServer(connList[i], i+1, j, conf.ReadFromFile)
			}

		}()
	}
	wg.Wait()

	server.CloseServer()
}

func main() {
	var server *src.Server
	src.CreateServer(&server)

	go src.StartListening(&server)

	if conf.ReadFromFile {
		handleExample(server)
	} else {
		<-server.CloseChan
		server.CloseServer()
	}
}
