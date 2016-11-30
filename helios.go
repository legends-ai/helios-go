package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/asunaio/helios/config"
	apb "github.com/asunaio/helios/gen-go/asuna"
	"github.com/asunaio/helios/server"
)

func main() {
	cfg := config.Initialize()
	if cfg == nil {
		log.Fatalf("Helios failed to initialize config.")
	}

	log.Printf("Connecting to Lucinda at %v", cfg.LucindaHost)
	conn, err := grpc.Dial(cfg.LucindaHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Lucinda: %v", err)
	}

	h := &server.Handlers{
		Lucinda: apb.NewLucindaClient(conn),
		Context: context.Background(),
	}

	go server.Monitor(cfg)
	server.Router(cfg, h)
}
