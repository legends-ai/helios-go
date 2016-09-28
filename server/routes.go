package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asunaio/helios/config"
)

func Monitor(cfg *config.AppConfig) {
	monitorMux := http.NewServeMux()
	monitorMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})
	monitorPort := fmt.Sprintf(":%d", cfg.MonitorPort)
	log.Printf("Monitor listening on port %d", cfg.MonitorPort)
	log.Fatal(http.ListenAndServe(monitorPort, monitorMux))
}

func Router(c *Controllers) {
	mux := http.NewServeMux()
	mux.HandleFunc("/champion", c.HandleChampion)
	mux.HandleFunc("/matchup", c.HandleMatchup)
}
