package src

import (
	e "client-server/config"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var CURRENT_ID = 1
var mapLock = &sync.Mutex{}

var lock = &sync.Mutex{}

type Server struct {
	ln          net.Listener
	CloseChan   chan struct{}
	request_map map[string]func([]string) string
	ClientTable map[string]string
	semaphore   chan struct{}
}

func CreateServer(server **Server) {
	if *server == nil {
		// Listen for incoming connections on port 8080
		ln, err := net.Listen("tcp", ":8080")
		e.PrintError(err)

		lock.Lock()
		defer lock.Unlock()
		if *server == nil {
			fmt.Println("Creating server instance now.")
			*server = &Server{ln: ln, CloseChan: make(chan struct{}), request_map: REQUEST_MAP, ClientTable: make(map[string]string), semaphore: make(chan struct{}, conf.ClientsNumber)}
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
		fmt.Printf("Client %s connected to the server!\n", (*server).GetClientId(conn, true))

		// _, err = conn.Write([]byte(fmt.Sprintf("NOTICE: Server accepts arrays of length <= %d\n", conf.ArrayLength)))
		e.PrintError(err)
		// first time handeling this client, we add it to the semaphore
		// since the semaphore channel is buffered, when it is full it will block any new connections, until it is freed
		(*server).semaphore <- struct{}{}

		go (*server).HandleConnection2(conn)

	}
}

func (server *Server) HandleConnection2(conn net.Conn) {
	// Read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	e.PrintError(err)

	input := strings.TrimSpace(string(buf[:n]))
	parts := strings.Split(input, " ")
	req := parts[0]

	if req == "exit" {
		// Close connection if client requested
		server.CloseConnection(conn)
		return
	}

	if len(parts) < 2 {
		_, _ = conn.Write([]byte("Invalid request\n"))
		server.HandleConnection2(conn)
		return
	}

	rows, _ := strconv.Atoi(parts[1])
	var args []string

	if rows == 1 {
		n, err = conn.Read(buf)
		e.PrintError(err)
		args = strings.Split(strings.TrimSpace(string(buf[:n])), " ")
		_, _ = conn.Write([]byte("Array received\n"))
	} else {
		for i := 0; i < rows; i++ {
			n, err := conn.Read(buf)
			e.PrintError(err)
			row := strings.Split(strings.TrimSpace(string(buf[:n])), " ")

			args = append(args, strings.Join(row, "; "))
		}
		_, _ = conn.Write([]byte("Matrix received\n"))
	}
	// Print the incoming data
	fmt.Printf("S: Received data: %v\n", args)

	// Process request
	responseFunc, exists := server.request_map[req]
	var response string
	if exists {
		response = responseFunc(args)
	} else {
		response = "Invalid request"
	}
	clientId := server.GetClientId(conn, true)
	fmt.Printf("Client %s requested: %s\n", clientId, req)

	// Send response back to client
	_, err = conn.Write([]byte(fmt.Sprintf("Message from Server: %s\n", response)))
	e.PrintError(err)

	server.HandleConnection2(conn)
}

func (server *Server) HandleConnection(conn net.Conn) {
	// Read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	e.PrintError(err)

	// Send a message to the client acknowledging the request
	_, err = conn.Write([]byte("Acknowledged request, processing input ...\n"))
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
	os.Exit(0)
}

func (server *Server) CloseConnection(conn net.Conn) {
	fmt.Printf("S: Closing connection with client %s\n", server.GetClientId(conn, true))
	server.DeleteClientId(conn)
	<-server.semaphore //release a client from the semaphore
	conn.Close()
}
