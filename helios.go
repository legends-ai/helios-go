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
	} else if cfg.ApolloHost == "" {
		log.Fatalf("Apollo Host not passed in via command line arguments.")
	}

	log.Printf("Connecting to Apollo at %v", cfg.ApolloHost)
	conn, err := grpc.Dial(cfg.ApolloHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Apollo: %v", err)
	}

	c := &server.Handlers{
		Apollo:  apb.NewApolloClient(conn),
		Context: context.Background(),
	}

	go server.Monitor(cfg)
	server.Router(cfg, c)
}
