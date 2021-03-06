// Code generated by counterfeiter. DO NOT EDIT.
package lockfakes

import (
	"sync"

	"github.com/chenbh/concourse/atc/db/lock"
)

type FakeLock struct {
	ReleaseStub        func() error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
	}
	releaseReturns struct {
		result1 error
	}
	releaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLock) Release() error {
	fake.releaseMutex.Lock()
	ret, specificReturn := fake.releaseReturnsOnCall[len(fake.releaseArgsForCall)]
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
	}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.releaseReturns
	return fakeReturns.result1
}

func (fake *FakeLock) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLock) ReleaseCalls(stub func() error) {
	fake.releaseMutex.Lock()
	defer fake.releaseMutex.Unlock()
	fake.ReleaseStub = stub
}

func (fake *FakeLock) ReleaseReturns(result1 error) {
	fake.releaseMutex.Lock()
	defer fake.releaseMutex.Unlock()
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLock) ReleaseReturnsOnCall(i int, result1 error) {
	fake.releaseMutex.Lock()
	defer fake.releaseMutex.Unlock()
	fake.ReleaseStub = nil
	if fake.releaseReturnsOnCall == nil {
		fake.releaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.releaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLock) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLock) recordInvocation(key string, args []interface{}) {
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

var _ lock.Lock = new(FakeLock)
