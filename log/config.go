package fxlog

import (
	"io"
	"os"

	"github.com/go-kit/kit/log/level"
)

var (
	// defaultOutput is the default io.Writer where log messages are written (os.Stdout).
	defaultOutput = os.Stdout

	// defaultFormat is the default log format (json)
	// Available formats: json, logfmt
	defaultFormat = JsonFormat

	// defaultFallbackLevel is the default fallback level used in messages when a level is not defined.
	defaultFallbackLevel level.Value = level.InfoValue()
)

// Config holds a list of options used during the logger construction.
type Config struct {
	Output        io.Writer
	Format        format
	FallbackLevel level.Value
	Context       []interface{}
	Debug         bool
}

// NewConfig returns a new config populated with default values.
func NewConfig() *Config {
	return &Config{
		Output:        defaultOutput,
		Format:        defaultFormat,
		FallbackLevel: defaultFallbackLevel,
	}
}
