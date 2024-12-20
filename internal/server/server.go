// Package server contains the server object and the NewServer function that creates a new server.
package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	// Load environment variables from .env file
	_ "github.com/joho/godotenv/autoload"
)

// Server object
type Server struct {
	port int
}

// NewServer creates a new server
func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
