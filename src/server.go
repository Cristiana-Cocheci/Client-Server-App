package src

import (
	e "client-server/config"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var request_map = map[string]func([]string) string{
	"Blank": func([]string) string { return "" },
	"ex1": func(args []string) string {
		return MixedLetters(args)
	},
	"ex3": func(args []string) string {
		return SolveEx3(args)
	},
	"ex5": func(args []string) string {
		return SolveEx5(args)
	},
	"ex6": func(args []string) string {
		return SolveEx6(args)
	},
	"ex7": func(args []string) string {
		return SolveEx7(args[0])
	},
	"ex8": func(args []string) string {
		return SolveEx8(args)
	},
	"ex12": func(args []string) string {
		return SolveEx12(args)
	},
	"exit": func([]string) string { return "exit" },
}

var lock = &sync.Mutex{}

type Server struct {
	ln net.Listener
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
			*server = &Server{ln: ln}
		} else {
			fmt.Println("Server already created.")
		}
	}
}

func StartListening(server **Server) {
	for {
		conn, err := (*server).ln.Accept()
		e.PrintError(err)

		go HandleConnection(conn)

	}
}

func HandleConnection(conn net.Conn) {
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
	responseFunc, exists := request_map[req]
	var response string
	if exists {
		response = responseFunc(args)
	} else {
		response = "Invalid request"
	}

	// Send response back to client
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client %s requested: %s\n", clientAddr, req)

	// Send response back to client
	_, err = conn.Write([]byte(fmt.Sprintf("Message from Server: %s\n", response)))
	e.PrintError(err)

	// Close connection if client requested
	if response == "exit" {
		CloseConnection(conn)
	} else {
		HandleConnection(conn)
	}
}

func CloseServer(server *Server) {
	if server == nil {
		fmt.Println("Server is not running.")
		return
	}
	fmt.Println("Closing server.")
	// server.ln.Close()
	os.Exit(0)
}

func CloseConnection(conn net.Conn) {
	fmt.Printf("S: Closing connection with client %s\n", conn.RemoteAddr().String())
	conn.Close()
}
