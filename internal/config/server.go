package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Port string
}

func NewServerConfig() *ServerConfig {
	checkServerEnv()
	return &ServerConfig{
		Port: os.Getenv("PORT"),
	}
}

func checkServerEnv() {
	envs := []string{"PORT"}
	for _, env := range envs {
		if os.Getenv(env) == "" {
			log.Fatalf("ENV[%s] required", env)
		}
	}
}
