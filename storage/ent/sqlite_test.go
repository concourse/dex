package ent

import (
	"log/slog"
	"testing"

	"github.com/concourse/dex/storage"
	"github.com/concourse/dex/storage/conformance"
)

func newSQLiteStorage(t *testing.T) storage.Storage {
	logger := slog.New(slog.NewTextHandler(t.Output(), &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg := SQLite3{File: ":memory:"}
	s, err := cfg.Open(logger)
	if err != nil {
		panic(err)
	}
	return s
}

func TestSQLite3(t *testing.T) {
	conformance.RunTests(t, newSQLiteStorage)
}
