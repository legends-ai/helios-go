package config

import "flag"

type AppConfig struct {
	ApolloHost  string
	Port        int
	MonitorPort int
}

func Initialize() *AppConfig {
	flag.Parse()
	return &AppConfig{
		ApolloHost:  *flag.String("apollo_host", "", "Apollo Host"),
		Port:        *flag.Int("port", 7921, "Host Port"),
		MonitorPort: *flag.Int("monitor_port", 7922, "Monitoring Port"),
	}
}
