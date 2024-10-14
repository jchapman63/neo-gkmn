package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jchapman63/neo-gkmn/internal/config"
)

//go:embed scripts/*.sql
var seed embed.FS

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("unable to load config", "err", err)
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.Database.URL)
	if err != nil {
		slog.Error("unable to connect to db", "err", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	scripts, err := seed.ReadDir("scripts")
	if err != nil {
		slog.Error("unable to read embedded directory", "err", err)
		os.Exit(1)
	}

	for _, script := range scripts {
		filePath := "scripts/" + script.Name()
		content, err := fs.ReadFile(seed, filePath)
		if err != nil {
			slog.Error("unable to read file contents", "err", err)
		}

		if _, err := conn.Exec(ctx, string(content)); err != nil {
			msg := fmt.Sprintf("Failed to execute SQL from file %s", filePath)
			slog.Error(msg, "err", err)
			continue
		}

		fmt.Printf("Successfully executed SQL from file %s\n", filePath)
	}
}
