// Code generated by counterfeiter. DO NOT EDIT.
package enginefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/engine"
	"github.com/chenbh/concourse/atc/exec"
)

type FakeStepBuilder struct {
	BuildStepStub        func(lager.Logger, db.Build) (exec.Step, error)
	buildStepMutex       sync.RWMutex
	buildStepArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Build
	}
	buildStepReturns struct {
		result1 exec.Step
		result2 error
	}
	buildStepReturnsOnCall map[int]struct {
		result1 exec.Step
		result2 error
	}
	BuildStepErroredStub        func(lager.Logger, db.Build, error)
	buildStepErroredMutex       sync.RWMutex
	buildStepErroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Build
		arg3 error
	}
	CheckStepStub        func(lager.Logger, db.Check) (exec.Step, error)
	checkStepMutex       sync.RWMutex
	checkStepArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Check
	}
	checkStepReturns struct {
		result1 exec.Step
		result2 error
	}
	checkStepReturnsOnCall map[int]struct {
		result1 exec.Step
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStepBuilder) BuildStep(arg1 lager.Logger, arg2 db.Build) (exec.Step, error) {
	fake.buildStepMutex.Lock()
	ret, specificReturn := fake.buildStepReturnsOnCall[len(fake.buildStepArgsForCall)]
	fake.buildStepArgsForCall = append(fake.buildStepArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Build
	}{arg1, arg2})
	fake.recordInvocation("BuildStep", []interface{}{arg1, arg2})
	fake.buildStepMutex.Unlock()
	if fake.BuildStepStub != nil {
		return fake.BuildStepStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.buildStepReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStepBuilder) BuildStepCallCount() int {
	fake.buildStepMutex.RLock()
	defer fake.buildStepMutex.RUnlock()
	return len(fake.buildStepArgsForCall)
}

func (fake *FakeStepBuilder) BuildStepCalls(stub func(lager.Logger, db.Build) (exec.Step, error)) {
	fake.buildStepMutex.Lock()
	defer fake.buildStepMutex.Unlock()
	fake.BuildStepStub = stub
}

func (fake *FakeStepBuilder) BuildStepArgsForCall(i int) (lager.Logger, db.Build) {
	fake.buildStepMutex.RLock()
	defer fake.buildStepMutex.RUnlock()
	argsForCall := fake.buildStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStepBuilder) BuildStepReturns(result1 exec.Step, result2 error) {
	fake.buildStepMutex.Lock()
	defer fake.buildStepMutex.Unlock()
	fake.BuildStepStub = nil
	fake.buildStepReturns = struct {
		result1 exec.Step
		result2 error
	}{result1, result2}
}

func (fake *FakeStepBuilder) BuildStepReturnsOnCall(i int, result1 exec.Step, result2 error) {
	fake.buildStepMutex.Lock()
	defer fake.buildStepMutex.Unlock()
	fake.BuildStepStub = nil
	if fake.buildStepReturnsOnCall == nil {
		fake.buildStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
			result2 error
		})
	}
	fake.buildStepReturnsOnCall[i] = struct {
		result1 exec.Step
		result2 error
	}{result1, result2}
}

func (fake *FakeStepBuilder) BuildStepErrored(arg1 lager.Logger, arg2 db.Build, arg3 error) {
	fake.buildStepErroredMutex.Lock()
	fake.buildStepErroredArgsForCall = append(fake.buildStepErroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Build
		arg3 error
	}{arg1, arg2, arg3})
	fake.recordInvocation("BuildStepErrored", []interface{}{arg1, arg2, arg3})
	fake.buildStepErroredMutex.Unlock()
	if fake.BuildStepErroredStub != nil {
		fake.BuildStepErroredStub(arg1, arg2, arg3)
	}
}

func (fake *FakeStepBuilder) BuildStepErroredCallCount() int {
	fake.buildStepErroredMutex.RLock()
	defer fake.buildStepErroredMutex.RUnlock()
	return len(fake.buildStepErroredArgsForCall)
}

func (fake *FakeStepBuilder) BuildStepErroredCalls(stub func(lager.Logger, db.Build, error)) {
	fake.buildStepErroredMutex.Lock()
	defer fake.buildStepErroredMutex.Unlock()
	fake.BuildStepErroredStub = stub
}

func (fake *FakeStepBuilder) BuildStepErroredArgsForCall(i int) (lager.Logger, db.Build, error) {
	fake.buildStepErroredMutex.RLock()
	defer fake.buildStepErroredMutex.RUnlock()
	argsForCall := fake.buildStepErroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStepBuilder) CheckStep(arg1 lager.Logger, arg2 db.Check) (exec.Step, error) {
	fake.checkStepMutex.Lock()
	ret, specificReturn := fake.checkStepReturnsOnCall[len(fake.checkStepArgsForCall)]
	fake.checkStepArgsForCall = append(fake.checkStepArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Check
	}{arg1, arg2})
	fake.recordInvocation("CheckStep", []interface{}{arg1, arg2})
	fake.checkStepMutex.Unlock()
	if fake.CheckStepStub != nil {
		return fake.CheckStepStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkStepReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStepBuilder) CheckStepCallCount() int {
	fake.checkStepMutex.RLock()
	defer fake.checkStepMutex.RUnlock()
	return len(fake.checkStepArgsForCall)
}

func (fake *FakeStepBuilder) CheckStepCalls(stub func(lager.Logger, db.Check) (exec.Step, error)) {
	fake.checkStepMutex.Lock()
	defer fake.checkStepMutex.Unlock()
	fake.CheckStepStub = stub
}

func (fake *FakeStepBuilder) CheckStepArgsForCall(i int) (lager.Logger, db.Check) {
	fake.checkStepMutex.RLock()
	defer fake.checkStepMutex.RUnlock()
	argsForCall := fake.checkStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStepBuilder) CheckStepReturns(result1 exec.Step, result2 error) {
	fake.checkStepMutex.Lock()
	defer fake.checkStepMutex.Unlock()
	fake.CheckStepStub = nil
	fake.checkStepReturns = struct {
		result1 exec.Step
		result2 error
	}{result1, result2}
}

func (fake *FakeStepBuilder) CheckStepReturnsOnCall(i int, result1 exec.Step, result2 error) {
	fake.checkStepMutex.Lock()
	defer fake.checkStepMutex.Unlock()
	fake.CheckStepStub = nil
	if fake.checkStepReturnsOnCall == nil {
		fake.checkStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
			result2 error
		})
	}
	fake.checkStepReturnsOnCall[i] = struct {
		result1 exec.Step
		result2 error
	}{result1, result2}
}

func (fake *FakeStepBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildStepMutex.RLock()
	defer fake.buildStepMutex.RUnlock()
	fake.buildStepErroredMutex.RLock()
	defer fake.buildStepErroredMutex.RUnlock()
	fake.checkStepMutex.RLock()
	defer fake.checkStepMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStepBuilder) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ engine.StepBuilder = new(FakeStepBuilder)
