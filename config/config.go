package config

import "flag"

type AppConfig struct {
	ApolloHost  string
	Port        int
	MonitorPort int
}

var (
	apolloHost  = flag.String("apollo_host", "", "Apollo Host")
	port        = flag.Int("port", 7921, "Host Port")
	monitorPort = flag.Int("monitor_port", 7922, "Monitoring Port")
)

func Initialize() *AppConfig {
	flag.Parse()
	return &AppConfig{
		ApolloHost:  *apolloHost,
		Port:        *port,
		MonitorPort: *monitorPort,
	}
}
