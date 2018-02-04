package test

// Runner is the interface responsible for running tests implemented by testing.M.
type Runner interface {
	// Run executes the tests and returns with an exit code.
	Run() int
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
