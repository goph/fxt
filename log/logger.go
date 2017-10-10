package log

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// NewLogger returns a new Logger.
func NewLogger(config *Config) (log.Logger, error) {
	var logger log.Logger

	w := log.NewSyncWriter(config.Output)

	switch config.Format {
	case LogfmtFormat:
		logger = log.NewLogfmtLogger(w)

	case JsonFormat:
		logger = log.NewJSONLogger(w)

	default:
		return nil, fmt.Errorf("unsupported log format: %s", config.Format.String())
	}

	// Add default context
	if len(config.Context) > 0 {
		logger = log.With(logger, config.Context...)
	}

	// Fallback to Info level
	logger = level.NewInjector(logger, config.FallbackLevel)

	// Only log debug level messages if debug mode is turned on
	if config.Debug == false {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	return logger, nil
}
