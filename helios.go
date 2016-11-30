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
	lucindaConn, err := grpc.Dial(cfg.LucindaHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Lucinda: %v", err)
	}

	log.Printf("Connecting to Vulgate at %v", cfg.VulgateHost)
	vulgateConn, err := grpc.Dial(cfg.VulgateHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Vulgate: %v", err)
	}

	h := &server.Handlers{
		Lucinda: apb.NewLucindaClient(lucindaConn),
		Vulgate: apb.NewVulgateClient(vulgateConn),
		Context: context.Background(),
	}

	go server.Monitor(cfg)
	server.Router(cfg, h)
}
