package internal

import "os"

type Config struct {
	Server   *Server
	Postgres *Postgres
}

func LoadConfig() *Config {
	server := loadServer()
	postgres := loadPostgres()

	return &Config{
		Server:   server,
		Postgres: postgres,
	}
}

type Server struct {
	Port string
	Mod  string
}

func loadServer() *Server {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	mod := os.Getenv("API_MODE")
	if mod == "" {
		mod = "debug"
	}

	return &Server{port, mod}
}

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func loadPostgres() *Postgres {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "db"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "postgres"
	}

	return &Postgres{host, port, user, password, name}
}
