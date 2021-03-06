// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/resource"
	"github.com/chenbh/concourse/atc/runtime"
	"github.com/chenbh/concourse/atc/worker"
)

type FakeFetchSourceFactory struct {
	NewFetchSourceStub        func(lager.Logger, worker.Worker, db.ContainerOwner, db.UsedResourceCache, resource.Resource, atc.VersionedResourceTypes, worker.ContainerSpec, runtime.ProcessSpec, db.ContainerMetadata, worker.ImageFetchingDelegate) worker.FetchSource
	newFetchSourceMutex       sync.RWMutex
	newFetchSourceArgsForCall []struct {
		arg1  lager.Logger
		arg2  worker.Worker
		arg3  db.ContainerOwner
		arg4  db.UsedResourceCache
		arg5  resource.Resource
		arg6  atc.VersionedResourceTypes
		arg7  worker.ContainerSpec
		arg8  runtime.ProcessSpec
		arg9  db.ContainerMetadata
		arg10 worker.ImageFetchingDelegate
	}
	newFetchSourceReturns struct {
		result1 worker.FetchSource
	}
	newFetchSourceReturnsOnCall map[int]struct {
		result1 worker.FetchSource
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSourceFactory) NewFetchSource(arg1 lager.Logger, arg2 worker.Worker, arg3 db.ContainerOwner, arg4 db.UsedResourceCache, arg5 resource.Resource, arg6 atc.VersionedResourceTypes, arg7 worker.ContainerSpec, arg8 runtime.ProcessSpec, arg9 db.ContainerMetadata, arg10 worker.ImageFetchingDelegate) worker.FetchSource {
	fake.newFetchSourceMutex.Lock()
	ret, specificReturn := fake.newFetchSourceReturnsOnCall[len(fake.newFetchSourceArgsForCall)]
	fake.newFetchSourceArgsForCall = append(fake.newFetchSourceArgsForCall, struct {
		arg1  lager.Logger
		arg2  worker.Worker
		arg3  db.ContainerOwner
		arg4  db.UsedResourceCache
		arg5  resource.Resource
		arg6  atc.VersionedResourceTypes
		arg7  worker.ContainerSpec
		arg8  runtime.ProcessSpec
		arg9  db.ContainerMetadata
		arg10 worker.ImageFetchingDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.recordInvocation("NewFetchSource", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10})
	fake.newFetchSourceMutex.Unlock()
	if fake.NewFetchSourceStub != nil {
		return fake.NewFetchSourceStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newFetchSourceReturns
	return fakeReturns.result1
}

func (fake *FakeFetchSourceFactory) NewFetchSourceCallCount() int {
	fake.newFetchSourceMutex.RLock()
	defer fake.newFetchSourceMutex.RUnlock()
	return len(fake.newFetchSourceArgsForCall)
}

func (fake *FakeFetchSourceFactory) NewFetchSourceCalls(stub func(lager.Logger, worker.Worker, db.ContainerOwner, db.UsedResourceCache, resource.Resource, atc.VersionedResourceTypes, worker.ContainerSpec, runtime.ProcessSpec, db.ContainerMetadata, worker.ImageFetchingDelegate) worker.FetchSource) {
	fake.newFetchSourceMutex.Lock()
	defer fake.newFetchSourceMutex.Unlock()
	fake.NewFetchSourceStub = stub
}

func (fake *FakeFetchSourceFactory) NewFetchSourceArgsForCall(i int) (lager.Logger, worker.Worker, db.ContainerOwner, db.UsedResourceCache, resource.Resource, atc.VersionedResourceTypes, worker.ContainerSpec, runtime.ProcessSpec, db.ContainerMetadata, worker.ImageFetchingDelegate) {
	fake.newFetchSourceMutex.RLock()
	defer fake.newFetchSourceMutex.RUnlock()
	argsForCall := fake.newFetchSourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10
}

func (fake *FakeFetchSourceFactory) NewFetchSourceReturns(result1 worker.FetchSource) {
	fake.newFetchSourceMutex.Lock()
	defer fake.newFetchSourceMutex.Unlock()
	fake.NewFetchSourceStub = nil
	fake.newFetchSourceReturns = struct {
		result1 worker.FetchSource
	}{result1}
}

func (fake *FakeFetchSourceFactory) NewFetchSourceReturnsOnCall(i int, result1 worker.FetchSource) {
	fake.newFetchSourceMutex.Lock()
	defer fake.newFetchSourceMutex.Unlock()
	fake.NewFetchSourceStub = nil
	if fake.newFetchSourceReturnsOnCall == nil {
		fake.newFetchSourceReturnsOnCall = make(map[int]struct {
			result1 worker.FetchSource
		})
	}
	fake.newFetchSourceReturnsOnCall[i] = struct {
		result1 worker.FetchSource
	}{result1}
}

func (fake *FakeFetchSourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFetchSourceMutex.RLock()
	defer fake.newFetchSourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetchSourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.FetchSourceFactory = new(FakeFetchSourceFactory)
