package main

import (
	"log"
	"os"
	"rawhttp/internal/server"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Print("PORT is required to run")
		return
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Print("invalid PORT")
		return
	}
	server.StartServer(port)
}
