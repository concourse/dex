package memory

import (
	"log/slog"
	"testing"

	"github.com/concourse/dex/storage"
	"github.com/concourse/dex/storage/conformance"
)

func TestStorage(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)

	newStorage := func() storage.Storage {
		return New(logger)
	}
	conformance.RunTests(t, newStorage)
}
