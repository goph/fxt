package fxpromhttp_test

import (
	"testing"

	"github.com/goph/fxt/metrics/prometheus/promhttp"
	"github.com/goph/fxt/testing/fxtest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestModule(t *testing.T) {
	var handler fxpromhttp.Handler

	app := fxtest.New(
		t,
		fx.Provide(func() prometheus.Gatherer {
			return prometheus.NewRegistry()
		}),
		fxpromhttp.Module,
		fx.Populate(&handler),
	)

	app.RequireStart().RequireStop()

	assert.NotNil(t, handler)
}
