package fxtest_test

import (
	"context"
	"testing"

	"github.com/goph/fxt"
	"github.com/goph/fxt/testing/fxtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestAppImplementsCloserInterface(t *testing.T) {
	assert.Implements(t, (*fxt.Closer)(nil), new(fxtest.App))
}

func TestNew(t *testing.T) {
	var lifecycle fxt.Lifecycle
	app := fxtest.New(
		t,
		fx.Populate(&lifecycle),
	)

	app.Start(context.Background())
	app.Stop(context.Background())

	assert.NotNil(t, lifecycle)
}

func TestApp_Close(t *testing.T) {
	called := false

	app := fxtest.New(
		t,
		fx.Invoke(func(l fxt.Lifecycle) {
			l.Append(fxt.Hook{
				OnClose: func() error {
					called = true

					return nil
				},
			})
		}),
	)

	err := app.Close()
	require.NoError(t, err)
	assert.True(t, called)
}
