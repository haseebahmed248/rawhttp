// TCP Listener, accept loop

package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"rawhttp/internal/handler"
	"rawhttp/internal/request"
	"rawhttp/internal/router"
)

func handleConn(conn net.Conn, route *router.Router) {
	defer conn.Close() // close connection after completion
	reader := bufio.NewReader(conn)
	result, err := request.Parse(reader)
	if err != nil {
		log.Print("Error: ", err)
		return
	}
	log.Print(result)

	// Response
	res := route.Handle(result)
	res.Write(conn)

}

func StartServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Print(err)
	}
	log.Print("Listening on port ", port)
	route := router.NewRouter()
	route.GET("/hello", handler.HelloHandler)
	route.GET("/health", handler.HealthHandler)
	route.POST("/echo", handler.EchoHandler)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn, route)
	}
}
