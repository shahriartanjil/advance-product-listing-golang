package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load the env file", err)
		os.Exit(1)
	}

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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	configurations = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), //type casting
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
