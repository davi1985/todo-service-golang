package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
