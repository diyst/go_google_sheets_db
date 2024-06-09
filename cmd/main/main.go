package main

import (
	"log"
	"net"

	"go_google_sheets_db/internal/config"
	"go_google_sheets_db/internal/server"
	"go_google_sheets_db/internal/sheets"
)

func main() {
	cfg, err := config.LoadConfig("credentials.json")
	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}

	err = sheets.StartClient(cfg)
	if err != nil {
		log.Fatalf("Unable to create Sheets client: %v", err)
	}

	listener, err := net.Listen("tcp", ":5432")
	if err != nil {
		log.Fatalf("Failed to listen on port 5432: %v", err)
	}
	defer listener.Close()
	log.Println("Listening on port 5432")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go server.HandleConnection(conn)
	}
}
