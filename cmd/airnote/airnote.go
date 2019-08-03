package main

import (
	"log"

	"github.com/stobita/airnote/internal/server"
)

func main() {
	log.Fatal(server.Run())
}
