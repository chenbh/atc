// Code generated by counterfeiter. DO NOT EDIT.
package builderfakes

import (
	"sync"

	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/engine/builder"
	"github.com/chenbh/concourse/atc/exec"
	"github.com/chenbh/concourse/vars"
)

type FakeDelegateFactory struct {
	BuildStepDelegateStub        func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.BuildStepDelegate
	buildStepDelegateMutex       sync.RWMutex
	buildStepDelegateArgsForCall []struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}
	buildStepDelegateReturns struct {
		result1 exec.BuildStepDelegate
	}
	buildStepDelegateReturnsOnCall map[int]struct {
		result1 exec.BuildStepDelegate
	}
	CheckDelegateStub        func(db.Check, atc.PlanID, vars.CredVarsTracker) exec.CheckDelegate
	checkDelegateMutex       sync.RWMutex
	checkDelegateArgsForCall []struct {
		arg1 db.Check
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}
	checkDelegateReturns struct {
		result1 exec.CheckDelegate
	}
	checkDelegateReturnsOnCall map[int]struct {
		result1 exec.CheckDelegate
	}
	GetDelegateStub        func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.GetDelegate
	getDelegateMutex       sync.RWMutex
	getDelegateArgsForCall []struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}
	getDelegateReturns struct {
		result1 exec.GetDelegate
	}
	getDelegateReturnsOnCall map[int]struct {
		result1 exec.GetDelegate
	}
	PutDelegateStub        func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.PutDelegate
	putDelegateMutex       sync.RWMutex
	putDelegateArgsForCall []struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}
	putDelegateReturns struct {
		result1 exec.PutDelegate
	}
	putDelegateReturnsOnCall map[int]struct {
		result1 exec.PutDelegate
	}
	TaskDelegateStub        func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.TaskDelegate
	taskDelegateMutex       sync.RWMutex
	taskDelegateArgsForCall []struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}
	taskDelegateReturns struct {
		result1 exec.TaskDelegate
	}
	taskDelegateReturnsOnCall map[int]struct {
		result1 exec.TaskDelegate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDelegateFactory) BuildStepDelegate(arg1 db.Build, arg2 atc.PlanID, arg3 vars.CredVarsTracker) exec.BuildStepDelegate {
	fake.buildStepDelegateMutex.Lock()
	ret, specificReturn := fake.buildStepDelegateReturnsOnCall[len(fake.buildStepDelegateArgsForCall)]
	fake.buildStepDelegateArgsForCall = append(fake.buildStepDelegateArgsForCall, struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}{arg1, arg2, arg3})
	fake.recordInvocation("BuildStepDelegate", []interface{}{arg1, arg2, arg3})
	fake.buildStepDelegateMutex.Unlock()
	if fake.BuildStepDelegateStub != nil {
		return fake.BuildStepDelegateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildStepDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeDelegateFactory) BuildStepDelegateCallCount() int {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	return len(fake.buildStepDelegateArgsForCall)
}

func (fake *FakeDelegateFactory) BuildStepDelegateCalls(stub func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = stub
}

func (fake *FakeDelegateFactory) BuildStepDelegateArgsForCall(i int) (db.Build, atc.PlanID, vars.CredVarsTracker) {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	argsForCall := fake.buildStepDelegateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDelegateFactory) BuildStepDelegateReturns(result1 exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = nil
	fake.buildStepDelegateReturns = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) BuildStepDelegateReturnsOnCall(i int, result1 exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = nil
	if fake.buildStepDelegateReturnsOnCall == nil {
		fake.buildStepDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.BuildStepDelegate
		})
	}
	fake.buildStepDelegateReturnsOnCall[i] = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) CheckDelegate(arg1 db.Check, arg2 atc.PlanID, arg3 vars.CredVarsTracker) exec.CheckDelegate {
	fake.checkDelegateMutex.Lock()
	ret, specificReturn := fake.checkDelegateReturnsOnCall[len(fake.checkDelegateArgsForCall)]
	fake.checkDelegateArgsForCall = append(fake.checkDelegateArgsForCall, struct {
		arg1 db.Check
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}{arg1, arg2, arg3})
	fake.recordInvocation("CheckDelegate", []interface{}{arg1, arg2, arg3})
	fake.checkDelegateMutex.Unlock()
	if fake.CheckDelegateStub != nil {
		return fake.CheckDelegateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeDelegateFactory) CheckDelegateCallCount() int {
	fake.checkDelegateMutex.RLock()
	defer fake.checkDelegateMutex.RUnlock()
	return len(fake.checkDelegateArgsForCall)
}

func (fake *FakeDelegateFactory) CheckDelegateCalls(stub func(db.Check, atc.PlanID, vars.CredVarsTracker) exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = stub
}

func (fake *FakeDelegateFactory) CheckDelegateArgsForCall(i int) (db.Check, atc.PlanID, vars.CredVarsTracker) {
	fake.checkDelegateMutex.RLock()
	defer fake.checkDelegateMutex.RUnlock()
	argsForCall := fake.checkDelegateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDelegateFactory) CheckDelegateReturns(result1 exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = nil
	fake.checkDelegateReturns = struct {
		result1 exec.CheckDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) CheckDelegateReturnsOnCall(i int, result1 exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = nil
	if fake.checkDelegateReturnsOnCall == nil {
		fake.checkDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.CheckDelegate
		})
	}
	fake.checkDelegateReturnsOnCall[i] = struct {
		result1 exec.CheckDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) GetDelegate(arg1 db.Build, arg2 atc.PlanID, arg3 vars.CredVarsTracker) exec.GetDelegate {
	fake.getDelegateMutex.Lock()
	ret, specificReturn := fake.getDelegateReturnsOnCall[len(fake.getDelegateArgsForCall)]
	fake.getDelegateArgsForCall = append(fake.getDelegateArgsForCall, struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetDelegate", []interface{}{arg1, arg2, arg3})
	fake.getDelegateMutex.Unlock()
	if fake.GetDelegateStub != nil {
		return fake.GetDelegateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeDelegateFactory) GetDelegateCallCount() int {
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	return len(fake.getDelegateArgsForCall)
}

func (fake *FakeDelegateFactory) GetDelegateCalls(stub func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.GetDelegate) {
	fake.getDelegateMutex.Lock()
	defer fake.getDelegateMutex.Unlock()
	fake.GetDelegateStub = stub
}

func (fake *FakeDelegateFactory) GetDelegateArgsForCall(i int) (db.Build, atc.PlanID, vars.CredVarsTracker) {
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	argsForCall := fake.getDelegateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDelegateFactory) GetDelegateReturns(result1 exec.GetDelegate) {
	fake.getDelegateMutex.Lock()
	defer fake.getDelegateMutex.Unlock()
	fake.GetDelegateStub = nil
	fake.getDelegateReturns = struct {
		result1 exec.GetDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) GetDelegateReturnsOnCall(i int, result1 exec.GetDelegate) {
	fake.getDelegateMutex.Lock()
	defer fake.getDelegateMutex.Unlock()
	fake.GetDelegateStub = nil
	if fake.getDelegateReturnsOnCall == nil {
		fake.getDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.GetDelegate
		})
	}
	fake.getDelegateReturnsOnCall[i] = struct {
		result1 exec.GetDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) PutDelegate(arg1 db.Build, arg2 atc.PlanID, arg3 vars.CredVarsTracker) exec.PutDelegate {
	fake.putDelegateMutex.Lock()
	ret, specificReturn := fake.putDelegateReturnsOnCall[len(fake.putDelegateArgsForCall)]
	fake.putDelegateArgsForCall = append(fake.putDelegateArgsForCall, struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}{arg1, arg2, arg3})
	fake.recordInvocation("PutDelegate", []interface{}{arg1, arg2, arg3})
	fake.putDelegateMutex.Unlock()
	if fake.PutDelegateStub != nil {
		return fake.PutDelegateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeDelegateFactory) PutDelegateCallCount() int {
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	return len(fake.putDelegateArgsForCall)
}

func (fake *FakeDelegateFactory) PutDelegateCalls(stub func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.PutDelegate) {
	fake.putDelegateMutex.Lock()
	defer fake.putDelegateMutex.Unlock()
	fake.PutDelegateStub = stub
}

func (fake *FakeDelegateFactory) PutDelegateArgsForCall(i int) (db.Build, atc.PlanID, vars.CredVarsTracker) {
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	argsForCall := fake.putDelegateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDelegateFactory) PutDelegateReturns(result1 exec.PutDelegate) {
	fake.putDelegateMutex.Lock()
	defer fake.putDelegateMutex.Unlock()
	fake.PutDelegateStub = nil
	fake.putDelegateReturns = struct {
		result1 exec.PutDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) PutDelegateReturnsOnCall(i int, result1 exec.PutDelegate) {
	fake.putDelegateMutex.Lock()
	defer fake.putDelegateMutex.Unlock()
	fake.PutDelegateStub = nil
	if fake.putDelegateReturnsOnCall == nil {
		fake.putDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.PutDelegate
		})
	}
	fake.putDelegateReturnsOnCall[i] = struct {
		result1 exec.PutDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) TaskDelegate(arg1 db.Build, arg2 atc.PlanID, arg3 vars.CredVarsTracker) exec.TaskDelegate {
	fake.taskDelegateMutex.Lock()
	ret, specificReturn := fake.taskDelegateReturnsOnCall[len(fake.taskDelegateArgsForCall)]
	fake.taskDelegateArgsForCall = append(fake.taskDelegateArgsForCall, struct {
		arg1 db.Build
		arg2 atc.PlanID
		arg3 vars.CredVarsTracker
	}{arg1, arg2, arg3})
	fake.recordInvocation("TaskDelegate", []interface{}{arg1, arg2, arg3})
	fake.taskDelegateMutex.Unlock()
	if fake.TaskDelegateStub != nil {
		return fake.TaskDelegateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.taskDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeDelegateFactory) TaskDelegateCallCount() int {
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	return len(fake.taskDelegateArgsForCall)
}

func (fake *FakeDelegateFactory) TaskDelegateCalls(stub func(db.Build, atc.PlanID, vars.CredVarsTracker) exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = stub
}

func (fake *FakeDelegateFactory) TaskDelegateArgsForCall(i int) (db.Build, atc.PlanID, vars.CredVarsTracker) {
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	argsForCall := fake.taskDelegateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDelegateFactory) TaskDelegateReturns(result1 exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = nil
	fake.taskDelegateReturns = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) TaskDelegateReturnsOnCall(i int, result1 exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = nil
	if fake.taskDelegateReturnsOnCall == nil {
		fake.taskDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.TaskDelegate
		})
	}
	fake.taskDelegateReturnsOnCall[i] = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeDelegateFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	fake.checkDelegateMutex.RLock()
	defer fake.checkDelegateMutex.RUnlock()
	fake.getDelegateMutex.RLock()
	defer fake.getDelegateMutex.RUnlock()
	fake.putDelegateMutex.RLock()
	defer fake.putDelegateMutex.RUnlock()
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDelegateFactory) recordInvocation(key string, args []interface{}) {
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

var _ builder.DelegateFactory = new(FakeDelegateFactory)