package test_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/goph/fxt/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGodogRunner_Run(t *testing.T) {
	buf := new(bytes.Buffer)

	runner := test.NewGodogRunner(test.WithGodogOptions(godog.Options{
		Output: buf,
	}))

	exitCode := runner.Run()
	require.Equal(t, 0, exitCode)
}

func TestGodogRunner_Run_Nil(t *testing.T) {
	var runner *test.GodogRunner

	exitCode := runner.Run()
	require.Equal(t, 3, exitCode)
}

func TestWithSuiteName(t *testing.T) {
	buf := new(bytes.Buffer)

	runner := test.NewGodogRunner(
		test.WithGodogOptions(godog.Options{
			Output: buf,
			Format: "events",
		}),
		test.WithSuiteName("suite_name"),
	)

	exitCode := runner.Run()
	require.Equal(t, 0, exitCode)

	rawEvents := strings.Split(buf.String(), "\n")
	if len(rawEvents) < 1 {
		t.Fatal("no events were recorded")
	}

	var event map[string]interface{}

	json.Unmarshal([]byte(rawEvents[0]), &event)

	if suiteName, ok := event["suite"].(string); ok {
		assert.Equal(t, "suite_name", suiteName)
	}
}

func TestGodogRunner_RegisterFeatureContext(t *testing.T) {
	buf := new(bytes.Buffer)

	runner := test.NewGodogRunner(test.WithGodogOptions(godog.Options{
		Output: buf,
	}))

	called := false

	runner.RegisterFeatureContext(func(s *godog.Suite) {
		called = true
	})

	exitCode := runner.Run()
	require.Equal(t, 0, exitCode)
	assert.True(t, called)
}

func TestGodogRunner_RegisterFeaturePath(t *testing.T) {
	buf := new(bytes.Buffer)

	runner := test.NewGodogRunner(test.WithGodogOptions(godog.Options{
		Output: buf,
	}))

	runner.RegisterFeaturePath("features")

	runner.RegisterFeatureContext(func(s *godog.Suite) {
		s.Step(`^I run Godog$`, func() error {
			return nil
		})
		s.Step(`^I should see test results$`, func() error {
			return nil
		})
	})

	exitCode := runner.Run()
	require.Equal(t, 0, exitCode)
}
