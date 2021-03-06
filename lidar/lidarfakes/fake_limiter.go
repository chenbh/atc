// Code generated by counterfeiter. DO NOT EDIT.
package lidarfakes

import (
	"context"
	"sync"

	"github.com/chenbh/concourse/atc/lidar"
)

type FakeLimiter struct {
	WaitStub        func(context.Context) error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		arg1 context.Context
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLimiter) Wait(arg1 context.Context) error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Wait", []interface{}{arg1})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1
}

func (fake *FakeLimiter) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeLimiter) WaitCalls(stub func(context.Context) error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeLimiter) WaitArgsForCall(i int) context.Context {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	argsForCall := fake.waitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLimiter) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLimiter) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLimiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLimiter) recordInvocation(key string, args []interface{}) {
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

var _ lidar.Limiter = new(FakeLimiter)
