package logging

import (
	"os"

	hclog "github.com/hashicorp/go-hclog"
)

// Logger is the interface that all logs are expected to use with this package.
type Logger interface {
	// Info logs a message with some additional context as key-value pairs
	Info(msg string, kvPairs ...interface{})
	// Error logs a message with some additional context as key-value pairs
	Error(msg string, kvPairs ...interface{})
	// Named adds a sub-scope to the logger's name
	Named(name string) Logger
	// LogIfError will log the info and error if the error is non-nil.
	// Returns true if err is non-nil.
	LogIfError(err error, kvPairs ...interface{}) bool
	// FatalIfError will log the error and exit if it is non-nil.
	// Note that calls to `defer` will not be triggered by this - no cleanup is done!
	FatalIfError(err error, kvPairs ...interface{})
}

type hclogWrapper struct {
	hclog.Logger
}

func (h *hclogWrapper) Named(name string) Logger {
	return &hclogWrapper{h.Logger.Named(name)}
}

func (h *hclogWrapper) LogIfError(err error, kvPairs ...interface{}) bool {
	if err != nil {
		h.Error("error", "error", err, kvPairs)
		return true
	}
	return false
}

func (h *hclogWrapper) FatalIfError(err error, kvPairs ...interface{}) {
	if err != nil {
		h.Error("FATAL error", "error", err, kvPairs)
		os.Exit(1)
	}
}

// NewLogger creates a new Logger, using hclog to back it for structured logging.
func NewLogger(name string) Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  name,
		Level: hclog.Info,
	})

	return &hclogWrapper{logger}
}
