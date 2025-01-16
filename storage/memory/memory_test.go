package memory

import (
	"io"
	"log/slog"
	"testing"

	"github.com/concourse/dex/storage"
	"github.com/concourse/dex/storage/conformance"
)

func TestStorage(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))

	newStorage := func() storage.Storage {
		return New(logger)
	}
	conformance.RunTests(t, newStorage)
}
