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

const CLIENTS = 5

func handleExample(server *src.Server) {
	wg := sync.WaitGroup{}
	var connList [CLIENTS]net.Conn

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

func handleInteractiveInput(server *src.Server) {
	wg := sync.WaitGroup{}
	var connList [CLIENTS]net.Conn
	for {
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
				src.SendRequestToServer(connList[i], i+1, -1, conf.ReadFromFile)
			}()
		}
		wg.Wait()
	}
}

func main() {
	var server *src.Server
	src.StartServer(&server)
	go src.StartListening(&server)

	if conf.ReadFromFile {
		handleExample(server)
	} else {
		handleInteractiveInput(server)
	}
}
