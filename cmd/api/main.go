package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error")
	}

	client, database, err := db.Connect(cfg)
	fmt.Printf("DEBUG: Connecting to URI: [%s]\n", cfg.MongoURI)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}
	fmt.Println("Connecting to URI:", cfg.MongoURI)

	defer func() {

		if err := db.Disconnect(client); err != nil {
			log.Println("mongo disconnected:", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}
}
