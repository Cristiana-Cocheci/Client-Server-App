package main

import (
	"client-server/config"
	"client-server/src"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
)

var wd, _ = os.Getwd()
var config_path = filepath.Join(wd, "env.txt")

var conf config.Config = config.LoadConfig(config_path)

func handleExample(server *src.Server) {
	wg := sync.WaitGroup{}
	connList := make(map[int]net.Conn)

	wg.Add(int(conf.ClientsNumber))
	for i := 0; i < int(conf.ClientsNumber); i++ {
		go func() {
			fmt.Println("New connection.")
			defer wg.Done()
			connList[i] = src.ConnectToServer(i, conf.ReadFromFile)
		}()
	}
	wg.Wait()

	wg.Add(int(conf.ClientsNumber))
	for i := 0; i < int(conf.ClientsNumber); i++ {
		go func() {
			defer wg.Done()
			req_number := src.GetReqNumber(i + 1)
			for j := 0; j < req_number; j++ {
				src.SendRequestToServer(connList[i], i+1, j, conf.ReadFromFile)
			}

		}()
	}
	wg.Wait()

	src.CloseServer(server)
}

func main() {
	var server *src.Server
	src.StartServer(&server)

	go src.StartListening(&server)

	if conf.ReadFromFile {
		handleExample(server)
	} else {
		<-server.CloseChan
		src.CloseServer(server)
	}
}
