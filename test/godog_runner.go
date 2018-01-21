package test

import (
	"flag"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/DATA-DOG/godog"
)

// GodogRunnerOption sets an option on the Godog runner.
type GodogRunnerOption func(*GodogRunner)

// WithSuiteName sets Godog suite name on the runner.
func WithSuiteName (s string) GodogRunnerOption {
	return func(r *GodogRunner) {
		r.suite = s
	}
}

// WithGodogOptions sets Godog options on the runner.
func WithGodogOptions (o godog.Options) GodogRunnerOption {
	return func(r *GodogRunner) {
		r.options = o
	}
}

// GodogRunner acts as a wrapper around Godog allowing to maintain a global acceptance test state.
type GodogRunner struct {
	featureContexts []func(s *godog.Suite)

	suite string
	options godog.Options

	mu sync.Mutex
}

// NewGodogRunner returns a new Godog runner.
func NewGodogRunner(opts ...GodogRunnerOption) *GodogRunner {
	runner := new(GodogRunner)

	for _, opt := range opts {
		opt(runner)
	}

	// Default suite name
	if runner.suite == "" {
		runner.suite = "godog"
	}

	return runner
}

// RegisterFeaturePath registers a path where .feature files can be found.
func (r *GodogRunner) RegisterFeaturePath(featurePath string) {
	r.registerFeaturePath(featurePath, 1)
}

func (r *GodogRunner) registerFeaturePath(featurePath string, callerDepth int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if path.IsAbs(featurePath) == false {
		_, filename, _, ok := runtime.Caller(callerDepth + 1)
		if !ok {
			panic("cannot determine feature path: no caller information")
		}

		featurePath = path.Clean(path.Join(path.Dir(filename), featurePath))
	}

	r.options.Paths = append(r.options.Paths, featurePath)
}

// RegisterFeatureContext registers a feature context
func (r *GodogRunner) RegisterFeatureContext(ctx func(s *godog.Suite)) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.featureContexts = append(r.featureContexts, ctx)
}

func (r *GodogRunner) Run() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	options := r.options

	if options.Format == "" {
		options.Format = "progress"

		// go test transforms -v option
		if verbose := flag.Lookup("test.v"); verbose != nil {
			options.Format = "pretty"
		}
	}

	if options.Randomize == 0 {
		// Randomize scenario execution order
		if randomize, _ := strconv.ParseBool(os.Getenv("TEST_RANDOMIZE")); randomize {
			options.Randomize = time.Now().UTC().UnixNano()
		}
	}

	return godog.RunWithOptions(
		r.suite,
		func(s *godog.Suite) {
			for _, featureContext := range r.featureContexts {
				featureContext(s)
			}
		},
		options,
	)
}
