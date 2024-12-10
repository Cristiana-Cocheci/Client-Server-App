package src

import (
	e "client-server/config"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var CURRENT_ID = 1
var mapLock = &sync.Mutex{}

var lock = &sync.Mutex{}

type Server struct {
	ln          net.Listener
	CloseChan   chan bool
	request_map map[string]func([]string) string
	ClientTable map[string]string
}

func StartServer(server **Server) {
	if *server == nil {
		// Listen for incoming connections on port 8080
		ln, err := net.Listen("tcp", ":8080")
		e.PrintError(err)

		lock.Lock()
		defer lock.Unlock()
		if *server == nil {
			fmt.Println("Creating server instance now.")
			*server = &Server{ln: ln, CloseChan: make(chan bool), request_map: REQUEST_MAP, ClientTable: make(map[string]string)}
		} else {
			fmt.Println("Server already created.")
		}
	}
}

func (server *Server) GetClientId(conn net.Conn, fromServer bool) string {
	mapLock.Lock()
	defer mapLock.Unlock()
	var connStr string
	if fromServer {
		connStr = conn.RemoteAddr().String()
	} else {
		connStr = conn.LocalAddr().String()
	}
	id, found := server.ClientTable[connStr]
	if found {
		return id
	} else {
		server.ClientTable[connStr] = fmt.Sprint(CURRENT_ID)
		CURRENT_ID++
		return server.ClientTable[connStr]
	}
}

func (server *Server) DeleteClientId(conn net.Conn) {
	mapLock.Lock()
	defer mapLock.Unlock()
	connStr := conn.RemoteAddr().String()
	delete(server.ClientTable, connStr)
}

func StartListening(server **Server) {
	for {
		conn, err := (*server).ln.Accept()
		e.PrintError(err)

		go (*server).HandleConnection(conn)

	}
}

func (server *Server) HandleConnection(conn net.Conn) {
	// Read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	e.PrintError(err)

	// Convert the buffer to string and process the request
	input := strings.TrimSpace(string(buf[:n]))
	parts := strings.Split(input, " ")
	req := parts[0]
	args := parts[1:]

	// Print the incoming data
	fmt.Printf("S: Received data: %s\n", buf)

	// Process request
	responseFunc, exists := server.request_map[req]
	var response string
	if exists {
		response = responseFunc(args)
	} else {
		response = "Invalid request"
	}

	// Send response back to client
	clientId := server.GetClientId(conn, true)
	fmt.Printf("Client %s requested: %s\n", clientId, req)

	// Send response back to client
	_, err = conn.Write([]byte(fmt.Sprintf("Message from Server: %s\n", response)))
	e.PrintError(err)

	// Close connection if client requested
	if response == "exit" {
		server.CloseConnection(conn)
	} else {
		server.HandleConnection(conn)
	}
}

func (server *Server) CloseServer() {
	if server == nil {
		fmt.Println("Server is not running.")
		return
	}
	fmt.Println("Closing server.")
	// server.ln.Close()
	os.Exit(0)
}

func (server *Server) CloseConnection(conn net.Conn) {
	fmt.Printf("S: Closing connection with client %s\n", server.GetClientId(conn, true))
	server.DeleteClientId(conn)
	conn.Close()
}
