package internal

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
	"task-management-system-api/internal/applications"
	"task-management-system-api/internal/infrastructure/users"
	"task-management-system-api/pkg/logger"
	"time"
)

func Run(cfg *Config) error {
	logger.Setup()

	db, err := postgresConnect(context.Background(), cfg)
	if err != nil {
		return err
	}

	slog.Info("server start ")
	if err := startServer(cfg.Server.Port, newHandler(db)); err != nil {
		return err
	}

	return nil
}

func startServer(port string, handler http.Handler) error {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func postgresConnect(ctx context.Context, cfg *Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database,
	)
	fmt.Println("Connection string:", connStr)

	fmt.Printf("%+v\n", cfg.Postgres)
	fmt.Println(os.Getenv("DB_USERNAME"))

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database,
	))
	if err != nil {
		slog.Error("unable to create connection pool: %v\n", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		slog.Error("unable to ping DB: %v\n", err)
		return nil, err
	}

	return pool, nil
}

func newHandler(db *pgxpool.Pool) http.Handler {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	user := users.NewPostgresRepo(db, qb)

	handler := applications.NewHTTPServer(user)
	return handler.SetupHTTPServer()
}
