package fxtesting_test

import (
	"testing"

	"github.com/goph/fxt/testing"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type RunnerMock struct {
	mock.Mock
}

func (_m *RunnerMock) Run() int {
	ret := _m.Called()

	return ret.Int(0)
}

type RunnerFactoryMock struct {
	mock.Mock
}

func (_m *RunnerFactoryMock) CreateRunner() (fxtesting.Runner, error) {
	ret := _m.Called()

	var r0 fxtesting.Runner
	if rf, ok := ret.Get(0).(func() fxtesting.Runner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fxtesting.Runner)
		}
	}

	return r0, ret.Error(1)
}

func TestRunners_Run(t *testing.T) {
	r := new(RunnerMock)
	r.On("Run").Return(0)

	runner := fxtesting.Runners{r}

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

	runner := fxtesting.Runners{r1, r2, r3}

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

	runner := fxtesting.AppendRunner(r1, r2, r3)

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

	runners := fxtesting.Runners{r1}
	runner := fxtesting.AppendRunner(runners, r2, r3)

	result := runner.Run()
	assert.Equal(t, 2, result)
	r1.AssertExpectations(t)
	r2.AssertExpectations(t)
	r3.AssertExpectations(t)

	runners, ok := runner.(fxtesting.Runners)
	require.True(t, ok)
	assert.Equal(t, r1, runners[0])
}

func TestAppendRunner_NothingToAppend(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	runner := fxtesting.AppendRunner(r1)

	result := runner.Run()
	assert.Equal(t, 0, result)
	assert.Equal(t, r1, runner)
	r1.AssertExpectations(t)
}

func TestRunnerFactoryRegistry(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	r2 := new(RunnerMock)
	r2.On("Run").Return(1)

	f1 := new(RunnerFactoryMock)
	f1.On("CreateRunner").Return(r1, nil)

	f2 := new(RunnerFactoryMock)
	f2.On("CreateRunner").Return(r2, nil)

	registry := fxtesting.RunnerFactoryRegistry{}

	registry.Register(f1)
	registry.Register(f2)

	runner, err := registry.CreateRunner()
	require.NoError(t, err)

	result := runner.Run()
	assert.Equal(t, 1, result)
	r1.AssertExpectations(t)
	r2.AssertExpectations(t)
	f1.AssertExpectations(t)
	f2.AssertExpectations(t)
}

func TestRunnerFactoryRegistry_Error(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	f1 := new(RunnerFactoryMock)
	f1.On("CreateRunner").Return(r1, nil)

	f2 := new(RunnerFactoryMock)
	f2err := errors.New("something went wrong")
	f2.On("CreateRunner").Return(nil, f2err)

	f3 := new(RunnerFactoryMock)

	registry := fxtesting.RunnerFactoryRegistry{}

	registry.Register(f1)
	registry.Register(f2)
	registry.Register(f3)

	runner, err := registry.CreateRunner()
	assert.Equal(t, f2err, err)
	assert.Nil(t, runner)

	f1.AssertExpectations(t)
	f2.AssertExpectations(t)
	f3.AssertNotCalled(t, "CreateRunner")
}

func TestRunnerFactoryFunc_CreateRunner(t *testing.T) {
	r1 := new(RunnerMock)
	r1.On("Run").Return(0)

	f1 := new(RunnerFactoryMock)
	f1.On("CreateRunner").Return(r1, nil)

	factoryFunc := fxtesting.RunnerFactoryFunc(f1.CreateRunner)

	runner, err := factoryFunc.CreateRunner()
	require.NoError(t, err)

	result := runner.Run()
	assert.Equal(t, 0, result)
	r1.AssertExpectations(t)
}
