// Code generated by counterfeiter. DO NOT EDIT.
package wrappafakes

import (
	"sync"

	"github.com/chenbh/concourse/atc/wrappa"
)

type FakePool struct {
	ReleaseStub        func()
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct {
	}
	SizeStub        func() int
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct {
	}
	sizeReturns struct {
		result1 int
	}
	sizeReturnsOnCall map[int]struct {
		result1 int
	}
	TryAcquireStub        func() bool
	tryAcquireMutex       sync.RWMutex
	tryAcquireArgsForCall []struct {
	}
	tryAcquireReturns struct {
		result1 bool
	}
	tryAcquireReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePool) Release() {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct {
	}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		fake.ReleaseStub()
	}
}

func (fake *FakePool) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakePool) ReleaseCalls(stub func()) {
	fake.releaseMutex.Lock()
	defer fake.releaseMutex.Unlock()
	fake.ReleaseStub = stub
}

func (fake *FakePool) Size() int {
	fake.sizeMutex.Lock()
	ret, specificReturn := fake.sizeReturnsOnCall[len(fake.sizeArgsForCall)]
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct {
	}{})
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sizeReturns
	return fakeReturns.result1
}

func (fake *FakePool) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakePool) SizeCalls(stub func() int) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = stub
}

func (fake *FakePool) SizeReturns(result1 int) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakePool) SizeReturnsOnCall(i int, result1 int) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	if fake.sizeReturnsOnCall == nil {
		fake.sizeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.sizeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakePool) TryAcquire() bool {
	fake.tryAcquireMutex.Lock()
	ret, specificReturn := fake.tryAcquireReturnsOnCall[len(fake.tryAcquireArgsForCall)]
	fake.tryAcquireArgsForCall = append(fake.tryAcquireArgsForCall, struct {
	}{})
	fake.recordInvocation("TryAcquire", []interface{}{})
	fake.tryAcquireMutex.Unlock()
	if fake.TryAcquireStub != nil {
		return fake.TryAcquireStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tryAcquireReturns
	return fakeReturns.result1
}

func (fake *FakePool) TryAcquireCallCount() int {
	fake.tryAcquireMutex.RLock()
	defer fake.tryAcquireMutex.RUnlock()
	return len(fake.tryAcquireArgsForCall)
}

func (fake *FakePool) TryAcquireCalls(stub func() bool) {
	fake.tryAcquireMutex.Lock()
	defer fake.tryAcquireMutex.Unlock()
	fake.TryAcquireStub = stub
}

func (fake *FakePool) TryAcquireReturns(result1 bool) {
	fake.tryAcquireMutex.Lock()
	defer fake.tryAcquireMutex.Unlock()
	fake.TryAcquireStub = nil
	fake.tryAcquireReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakePool) TryAcquireReturnsOnCall(i int, result1 bool) {
	fake.tryAcquireMutex.Lock()
	defer fake.tryAcquireMutex.Unlock()
	fake.TryAcquireStub = nil
	if fake.tryAcquireReturnsOnCall == nil {
		fake.tryAcquireReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.tryAcquireReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakePool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	fake.tryAcquireMutex.RLock()
	defer fake.tryAcquireMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePool) recordInvocation(key string, args []interface{}) {
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

var _ wrappa.Pool = new(FakePool)
