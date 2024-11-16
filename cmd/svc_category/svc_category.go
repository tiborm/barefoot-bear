package main

import (
	"log"
	"os"

	"github.com/tiborm/barefoot-bear/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    server.StartServer(logger)
}
