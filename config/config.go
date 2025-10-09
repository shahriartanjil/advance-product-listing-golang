package config

import (
	"fmt"
	"os"
	"strconv"
)

var config Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func loadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		// log.Fatal("Version is required")
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTO_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}
}

func GetConfig() Config {
	loadConfig()
	return config
}
