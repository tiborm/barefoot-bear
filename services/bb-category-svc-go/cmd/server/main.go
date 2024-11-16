package main

import (
	"log"
	"os"

	"github.com/tiborm/barefoot-bear/bb-catalog-svc-go/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    server.StartServer(logger)
}
