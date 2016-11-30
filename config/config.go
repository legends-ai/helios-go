package config

import "flag"

type AppConfig struct {
	LucindaHost string
	VulgateHost string
	Port        int
	MonitorPort int
}

var (
	lucindaHost = flag.String("lucinda_host", "localhost:45045", "Lucinda Host")
	vulgateHost = flag.String("vulgate_host", "localhost:6205", "Vulgate Host")
	port        = flag.Int("port", 7921, "Host Port")
	monitorPort = flag.Int("monitor_port", 7922, "Monitoring Port")
)

func Initialize() *AppConfig {
	flag.Parse()
	return &AppConfig{
		LucindaHost: *lucindaHost,
		VulgateHost: *vulgateHost,
		Port:        *port,
		MonitorPort: *monitorPort,
	}
}
