package main

import (
	"client-server/config"
	"client-server/src"
	"fmt"
	"net"
	"sync"
)

var config_path = "/Users/cricoche/Desktop/fmi/SD/client-server/env.txt"

var conf config.Config = config.LoadConfig(config_path)

const CLIENTS = 5

func main() {
	var server *src.Server
	src.StartServer(&server)
	go src.StartListening(&server)

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
				if conf.ReadFromFile {
					req_number := src.GetReqNumber(i + 1)
					for j := 0; j < req_number; j++ {
						src.SendRequestToServer(connList[i], i+1, j, conf.ReadFromFile)
					}
				} else {
					src.SendRequestToServer(connList[i], i+1, -1, conf.ReadFromFile)
				}
			}()
		}
		wg.Wait()
		if conf.ReadFromFile {
			break
		}
	}
	src.CloseServer(server)
}
