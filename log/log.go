package log

import (
	"errors"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

const LOG_LOCATION = "/var/log/tealok/events.json"

func NetworkCreated(name string) {
	slog.Info("Network created",
		"event", "network.created",
		"name", name,
	)
}

func Setup() error {
	// Create the log file
	file, err := os.OpenFile(LOG_LOCATION, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("Failed to open log file")
	}
	defer file.Close()

	// Create a logger that writes JSON to the file
	logger := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(file, &slog.HandlerOptions{}),
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}),
		),
	)

	slog.SetDefault(logger)

	// Log an example message with structured data
	slog.Info(
		"Tealok started",
		"event", "tealok.started",
	)
	return nil
}
