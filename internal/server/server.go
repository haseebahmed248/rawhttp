// TCP Listener, accept loop

package server
import (
 	"net"
	"log"
	"bufio"
	"fmt"
	"rawhttp/internal/request"
)

func handleConn(conn net.Conn){
	defer conn.Close()				// close connection after completion
	reader := bufio.NewReader(conn)
	result, err := request.Parse(reader)
	if err != nil{
		log.Print("Error: ", err)
		return
	}
	log.Print(result)
	
}

func StartServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Print(err)
	}
	log.Print("Listening on port ", port)
	for{
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}