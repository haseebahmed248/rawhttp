package main
import (
 	"net"
	"log"
	"bufio"
)

func handleConn(conn net.Conn){
	defer conn.Close()				// close connection after completion
	reader := bufio.NewReader(conn);
	for {
		status, err := reader.ReadString('\n')
		if err != nil{
			log.Print("Error: ", err)
			return;
		}
		log.Print(status)	
	}
	
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Print(err)
	}
	log.Print("Listening on port 9090")
	for{
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}