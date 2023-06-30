package ent

import (
	"io"
	"log/slog"
	"testing"

	"github.com/concourse/dex/storage"
	"github.com/concourse/dex/storage/conformance"
)

func newSQLiteStorage() storage.Storage {
	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))

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
