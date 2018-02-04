package test

import "sync"

// Runner is the interface responsible for running tests implemented by testing.M.
type Runner interface {
	// Run executes the tests and returns with an exit code.
	Run() int
}

// runnerFactory creates a new runner.
type runnerFactory interface {
	CreateRunner() (Runner, error)
}

// Runners is a list of test runners.
type Runners []Runner

// Run executes the underlying runners and returns the highest exit code.
func (r Runners) Run() int {
	var result int

	for _, runner := range r {
		r := runner.Run()
		if r > result {
			result = r
		}
	}

	return result
}

// AppendRunner checks if the target runner is already a runner list and appends the runner to it.
// Otherwise it creates a runner list and appends both runners to it.
func AppendRunner(target Runner, runners ...Runner) Runner {
	// Do not waste resources creating a runner list when there is nothing to append.
	if len(runners) == 0 {
		return target
	}

	r, ok := target.(Runners)
	if !ok {
		r = make(Runners, 1, len(runners)+1)
		r[0] = target
	}

	r = append(r, runners...)

	return r
}

// RunnerRegistry accepts runner factory implementations and creates a runner list from them.
type RunnerRegistry struct {
	factories []runnerFactory

	mu sync.Mutex
}

// Register appends a runner factory to the list.
// It is safe to call this method from multiple goroutines if necessary.
func (r *RunnerRegistry) Register(factory runnerFactory) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.factories = append(r.factories, factory)
}

// CreateRunner creates test runners from the underlying factories.
func (r *RunnerRegistry) CreateRunner() (Runner, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	runners := make(Runners, len(r.factories))

	for index, factory := range r.factories {
		runner, err := factory.CreateRunner()
		if err != nil {
			return nil, err
		}

		runners[index] = runner
	}

	return runners, nil
}

// RunnerFactoryFunc wraps a function implementing the runnerFactory interface.
type RunnerFactoryFunc func() (Runner, error)

func (fn RunnerFactoryFunc) CreateRunner() (Runner, error) {
	return fn()
}
