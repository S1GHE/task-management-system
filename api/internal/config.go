package internal

import "os"

type Config struct {
	Port string
	Mod  string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mod := os.Getenv("MOD")
	if mod == "" {
		mod = "dev"
	}

	return &Config{
		Port: port,
		Mod:  mod,
	}
}
