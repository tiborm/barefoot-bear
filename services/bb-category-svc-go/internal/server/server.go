package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tiborm/barefoot-bear/bb-catalog-svc-go/internal/middlewares"
	"github.com/tiborm/barefoot-bear/bb-catalog-svc-go/internal/routes"
)

func StartServer(
	logger *log.Logger,
	// Add stores and other services as dependencies like
	// config *oauth2.Config,
	// tenantsStore,
	// commentsStore,
	// conversationService,
	// chatGPTService,
) {
	mux := http.NewServeMux()
	
	routes.AddRoutes(mux)
	handler := middlewares.AddMiddlewares(mux)

	fmt.Println("Server started at http://localhost:5491")
	err := http.ListenAndServe(":5491", handler)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
