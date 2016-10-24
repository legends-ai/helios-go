package server

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/asunaio/helios/config"
)

func Monitor(cfg *config.AppConfig) {
	monitor := gin.New()
	monitor.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	monitor.Run(fmt.Sprintf(":%d", cfg.MonitorPort))
}

func Router(cfg *config.AppConfig, h *Handlers) {
	server := gin.New()
	server.GET("/champion/:id", h.HandleChampion)
	server.GET("/matchup/:focus/:enemy", h.HandleMatchup)
	server.Run(fmt.Sprintf(":%d", cfg.Port))
}
