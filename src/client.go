package src

import (
	"bufio"
	e "client-server/config"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func readRequest() (string, []string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	req := strings.Split(text, " ")[0]
	inputList := strings.Split(text, " ")[1:]
	if _, ok := REQUEST_MAP[req]; !ok {
		return text, []string{}
	}
	return req, inputList
}

func GetReqNumber(i int) int {
	var requests_path = filepath.Join(WD, "example_requests", fmt.Sprint(i)+".txt")
	return len(e.LoadRequests(requests_path))
}

func readRequestFromFile(i int, idx int) (string, []string) {
	var requests_path = filepath.Join(WD, "example_requests", fmt.Sprint(i)+".txt")
	s := strings.Split(e.LoadRequests(requests_path)[idx], " ")
	req := s[0]
	inputList := s[1:]
	if _, ok := REQUEST_MAP[req]; !ok {
		return req, []string{}
	}
	return req, inputList
}

func (server *Server) ConnectToServer(i int, readFromFile bool) net.Conn {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	e.PrintError(err)

	client_id := server.GetClientId(conn, false)
	fmt.Printf("Client %s connected to the server!\n", client_id)
	return conn
}

func (server *Server) SendRequestToServer(conn net.Conn, i int, idx int, readFromFile bool) {
	client_id := server.GetClientId(conn, false)
	var req string
	var inputList []string
	if readFromFile {
		req, inputList = readRequestFromFile(i, idx)
	} else {
		req, inputList = readRequest()
	}

	fmt.Printf("C %s: %s\n", client_id, req)

	// Combine request and arguments into a single message
	message := req + " " + strings.Join(inputList, " ")
	// Send request to the server
	_, err := conn.Write([]byte(message))
	e.PrintError(err)

	// Read the server's response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	e.PrintError(err)

	// Print the server's response
	fmt.Printf("C %s: %s\n", client_id, string(buf[:n]))

}
