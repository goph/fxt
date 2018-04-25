package logrus

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt/internal/keyvals"
	"github.com/sirupsen/logrus"
)

// Option sets options in the logrus adapter.
type Option func(*logrusAdapter)

// Logger sets a logrus instance in the logrus adapter.
func Logger(logger logrus.FieldLogger) func(*logrusAdapter) {
	return func(a *logrusAdapter) { a.logger = logger }
}

// MessageKey sets the key for the actual log message. By default, it's "msg".
func MessageKey(key string) Option {
	return func(a *logrusAdapter) { a.messageKey = key }
}

// LevelKey sets the key for the log level. By default, it's "level".
func LevelKey(key string) Option {
	return func(a *logrusAdapter) { a.levelKey = key }
}

// DefaultLevel sets the default log level. By default, it's INFO.
func DefaultLevel(lvl level.Value) Option {
	return func(a *logrusAdapter) { a.defaultLevel = lvl }
}

// logrusAdapter wraps a logrus logger and exposes it under a go-kit interface.
type logrusAdapter struct {
	logger logrus.FieldLogger

	messageKey   string
	levelKey     string
	defaultLevel level.Value
}

// New returns a new go-kit logger instance wrapping a logrus logger.
func New(options ...Option) log.Logger {
	adapter := &logrusAdapter{
		messageKey:   "msg",
		levelKey:     "level",
		defaultLevel: level.InfoValue(),
	}

	for _, o := range options {
		o(adapter)
	}

	// Default logrus instance
	if adapter.logger == nil {
		adapter.logger = logrus.New()
	}

	return adapter
}

func (a *logrusAdapter) Log(kv ...interface{}) error {
	kvmap := keyvals.ToMap(kv)

	var msg string

	if m, ok := kvmap[a.messageKey].(string); ok {
		msg = m

		delete(kvmap, a.messageKey)
	}

	var lvl interface{}
	lvl = a.defaultLevel

	if l, ok := kvmap[a.levelKey]; ok {
		lvl = l

		delete(kvmap, a.levelKey)
	}

	logger := a.logger.WithFields(logrus.Fields(kvmap))

	switch lvl {
	case level.DebugValue(), "debug":
		logger.Debug(msg)

	case level.InfoValue(), "info":
		logger.Info(msg)

	case level.WarnValue(), "warn":
		logger.Warn(msg)

	case level.ErrorValue(), "error":
		logger.Error(msg)

	default:
		logger.Print(msg)
	}

	return nil
}
