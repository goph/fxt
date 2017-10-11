package log

import (
	"fmt"
)

// format represents the supported log formats.
type format int

// String returns the format in a string representation.
func (f format) String() string {
	return formatMap[f]
}

// Format constants
const (
	JsonFormat format = iota
	LogfmtFormat
)

var (
	formatMap = map[format]string{
		JsonFormat:   "json",
		LogfmtFormat: "logfmt",
	}

	formatNameMap = map[string]format{
		"json":   JsonFormat,
		"logfmt": LogfmtFormat,
	}
)

// ParseFormat parses a string format name and returns the format or an error if the format is invalid.
func ParseFormat(formatName string) (format, error) {
	f, ok := formatNameMap[formatName]

	if !ok {
		return 0, fmt.Errorf("invalid log format: %s", formatName)
	}

	return f, nil
}
