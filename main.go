package main

import (
	"flag"
	"flight-tracker/domain"
	"flight-tracker/server"
	"fmt"
	"log"
)

func main() {
	var portFlag = flag.Int("port", 1234, "port of the server, default 1234")
	flag.Parse()
	if *portFlag == 0 {
		log.Fatal("port can not be 0")
	}
	sorter := domain.Domain{}
	server := server.NewServer(sorter)
	port := fmt.Sprintf(":%d", *portFlag)
	log.Println("Running on " + port)
	server.Run(port)
}
