package main

import (
	"github.com/SergeiSadov/positions/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err)
		return
	}
}
