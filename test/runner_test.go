package test_test

import (
	"testing"

	"github.com/goph/fxt/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type RunnerMock struct {
	mock.Mock
}

func (r *RunnerMock) Run() int {
	args := r.Called()

	return args.Int(0)
}

func TestRunners_Run(t *testing.T) {
	r := new(RunnerMock)
	r.On("Run").Return(0)

	runner := test.Runners{r}

	result := runner.Run()
	assert.Equal(t, 0, result)
	r.AssertExpectations(t)
}

func TestRunners_Run_Result(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	r2 := new(RunnerMock)
	r2.On("Run").Return(1)

	r3 := new(RunnerMock)
	r3.On("Run").Return(2)

	runner := test.Runners{r1, r2, r3}

	result := runner.Run()
	assert.Equal(t, 2, result)
	r1.AssertExpectations(t)
	r2.AssertExpectations(t)
	r3.AssertExpectations(t)
}

func TestAppendRunner(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	r2 := new(RunnerMock)
	r2.On("Run").Return(1)

	r3 := new(RunnerMock)
	r3.On("Run").Return(2)

	runner := test.AppendRunner(r1, r2, r3)

	result := runner.Run()
	assert.Equal(t, 2, result)
	r1.AssertExpectations(t)
	r2.AssertExpectations(t)
	r3.AssertExpectations(t)
}

func TestAppendRunner_Runners(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	r2 := new(RunnerMock)
	r2.On("Run").Return(1)

	r3 := new(RunnerMock)
	r3.On("Run").Return(2)

	runners := test.Runners{r1}
	runner := test.AppendRunner(runners, r2, r3)

	result := runner.Run()
	assert.Equal(t, 2, result)
	r1.AssertExpectations(t)
	r2.AssertExpectations(t)
	r3.AssertExpectations(t)

	runners, ok := runner.(test.Runners)
	require.True(t, ok)
	assert.Equal(t, r1, runners[0])
}

func TestAppendRunner_NothingToAppend(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	runner := test.AppendRunner(r1)

	result := runner.Run()
	assert.Equal(t, 0, result)
	assert.Equal(t, r1, runner)
	r1.AssertExpectations(t)
}
