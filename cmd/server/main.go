// @title           Todo API
// @version         1.0
// @description     A complete REST API for managing todos using Go and Gin
// @contact.name    API Support
// @host            localhost:8082
// @BasePath        /api/v1
// @schemes         http
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "todo-api/docs"
	"todo-api/internal/server"
)

func main() {
	log.Println("TODO API starting...")
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer srv.Close()
		
	go func() {
		if err := srv.Start(":8082"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	log.Println("Shutting down server...")
}
