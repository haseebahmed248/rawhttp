// TCP Listener, accept loop

package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"rawhttp/internal/request"
	"rawhttp/internal/response"
)

func handleConn(conn net.Conn) {
	defer conn.Close() // close connection after completion
	reader := bufio.NewReader(conn)
	result, err := request.Parse(reader)
	if err != nil {
		log.Print("Error: ", err)
		return
	}
	log.Print(result)

	// Response
	res := &response.Response{
		StatusCode: 200,
		StatusText: "OK",
		Body:       []byte("Hello World"),
	}
	res.Write(conn)

}

func StartServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Print(err)
	}
	log.Print("Listening on port ", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}
