package main

import (
	"guild-hack-api/app/interfaces/api"
	"log"
)

func main() {
	var port int
	s := api.NewServer()
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	port = 8080
	s.Run(port)
}
