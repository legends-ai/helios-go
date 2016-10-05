package server

import (
	"fmt"

	"github.com/kataras/iris"

	"github.com/asunaio/helios/config"
)

func Monitor(cfg *config.AppConfig) {
	monitor := iris.New()
	monitor.Get("/health", func(ctx *iris.Context) {
		ctx.Write("OK")
	})
	monitor.Listen(fmt.Sprintf(":%d", cfg.MonitorPort))
}

func Router(cfg *config.AppConfig, h *Handlers) {
	server := iris.New()
	server.Get("/champion/:id", h.HandleChampion)
	server.Get("/matchup/:focus/:enemy", h.HandleMatchup)
	server.Listen(fmt.Sprintf(":%d", cfg.Port))
}
