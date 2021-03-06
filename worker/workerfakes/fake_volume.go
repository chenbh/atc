// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/baggageclaim"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/worker"
)

type FakeVolume struct {
	COWStrategyStub        func() baggageclaim.COWStrategy
	cOWStrategyMutex       sync.RWMutex
	cOWStrategyArgsForCall []struct {
	}
	cOWStrategyReturns struct {
		result1 baggageclaim.COWStrategy
	}
	cOWStrategyReturnsOnCall map[int]struct {
		result1 baggageclaim.COWStrategy
	}
	CreateChildForContainerStub        func(db.CreatingContainer, string) (db.CreatingVolume, error)
	createChildForContainerMutex       sync.RWMutex
	createChildForContainerArgsForCall []struct {
		arg1 db.CreatingContainer
		arg2 string
	}
	createChildForContainerReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createChildForContainerReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	GetResourceCacheIDStub        func() int
	getResourceCacheIDMutex       sync.RWMutex
	getResourceCacheIDArgsForCall []struct {
	}
	getResourceCacheIDReturns struct {
		result1 int
	}
	getResourceCacheIDReturnsOnCall map[int]struct {
		result1 int
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	InitializeArtifactStub        func(string, int) (db.WorkerArtifact, error)
	initializeArtifactMutex       sync.RWMutex
	initializeArtifactArgsForCall []struct {
		arg1 string
		arg2 int
	}
	initializeArtifactReturns struct {
		result1 db.WorkerArtifact
		result2 error
	}
	initializeArtifactReturnsOnCall map[int]struct {
		result1 db.WorkerArtifact
		result2 error
	}
	InitializeResourceCacheStub        func(db.UsedResourceCache) error
	initializeResourceCacheMutex       sync.RWMutex
	initializeResourceCacheArgsForCall []struct {
		arg1 db.UsedResourceCache
	}
	initializeResourceCacheReturns struct {
		result1 error
	}
	initializeResourceCacheReturnsOnCall map[int]struct {
		result1 error
	}
	InitializeTaskCacheStub        func(lager.Logger, int, string, string, bool) error
	initializeTaskCacheMutex       sync.RWMutex
	initializeTaskCacheArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
		arg4 string
		arg5 bool
	}
	initializeTaskCacheReturns struct {
		result1 error
	}
	initializeTaskCacheReturnsOnCall map[int]struct {
		result1 error
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
	}
	pathReturns struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	PropertiesStub        func() (baggageclaim.VolumeProperties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
	}
	propertiesReturns struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	SetPrivilegedStub        func(bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(string, string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, baggageclaim.Encoding, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	WorkerNameStub        func() string
	workerNameMutex       sync.RWMutex
	workerNameArgsForCall []struct {
	}
	workerNameReturns struct {
		result1 string
	}
	workerNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolume) COWStrategy() baggageclaim.COWStrategy {
	fake.cOWStrategyMutex.Lock()
	ret, specificReturn := fake.cOWStrategyReturnsOnCall[len(fake.cOWStrategyArgsForCall)]
	fake.cOWStrategyArgsForCall = append(fake.cOWStrategyArgsForCall, struct {
	}{})
	fake.recordInvocation("COWStrategy", []interface{}{})
	fake.cOWStrategyMutex.Unlock()
	if fake.COWStrategyStub != nil {
		return fake.COWStrategyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cOWStrategyReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) COWStrategyCallCount() int {
	fake.cOWStrategyMutex.RLock()
	defer fake.cOWStrategyMutex.RUnlock()
	return len(fake.cOWStrategyArgsForCall)
}

func (fake *FakeVolume) COWStrategyCalls(stub func() baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = stub
}

func (fake *FakeVolume) COWStrategyReturns(result1 baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = nil
	fake.cOWStrategyReturns = struct {
		result1 baggageclaim.COWStrategy
	}{result1}
}

func (fake *FakeVolume) COWStrategyReturnsOnCall(i int, result1 baggageclaim.COWStrategy) {
	fake.cOWStrategyMutex.Lock()
	defer fake.cOWStrategyMutex.Unlock()
	fake.COWStrategyStub = nil
	if fake.cOWStrategyReturnsOnCall == nil {
		fake.cOWStrategyReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.COWStrategy
		})
	}
	fake.cOWStrategyReturnsOnCall[i] = struct {
		result1 baggageclaim.COWStrategy
	}{result1}
}

func (fake *FakeVolume) CreateChildForContainer(arg1 db.CreatingContainer, arg2 string) (db.CreatingVolume, error) {
	fake.createChildForContainerMutex.Lock()
	ret, specificReturn := fake.createChildForContainerReturnsOnCall[len(fake.createChildForContainerArgsForCall)]
	fake.createChildForContainerArgsForCall = append(fake.createChildForContainerArgsForCall, struct {
		arg1 db.CreatingContainer
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateChildForContainer", []interface{}{arg1, arg2})
	fake.createChildForContainerMutex.Unlock()
	if fake.CreateChildForContainerStub != nil {
		return fake.CreateChildForContainerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createChildForContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) CreateChildForContainerCallCount() int {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	return len(fake.createChildForContainerArgsForCall)
}

func (fake *FakeVolume) CreateChildForContainerCalls(stub func(db.CreatingContainer, string) (db.CreatingVolume, error)) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = stub
}

func (fake *FakeVolume) CreateChildForContainerArgsForCall(i int) (db.CreatingContainer, string) {
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	argsForCall := fake.createChildForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) CreateChildForContainerReturns(result1 db.CreatingVolume, result2 error) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = nil
	fake.createChildForContainerReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) CreateChildForContainerReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.createChildForContainerMutex.Lock()
	defer fake.createChildForContainerMutex.Unlock()
	fake.CreateChildForContainerStub = nil
	if fake.createChildForContainerReturnsOnCall == nil {
		fake.createChildForContainerReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createChildForContainerReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
	}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolume) DestroyCalls(stub func() error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolume) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) GetResourceCacheID() int {
	fake.getResourceCacheIDMutex.Lock()
	ret, specificReturn := fake.getResourceCacheIDReturnsOnCall[len(fake.getResourceCacheIDArgsForCall)]
	fake.getResourceCacheIDArgsForCall = append(fake.getResourceCacheIDArgsForCall, struct {
	}{})
	fake.recordInvocation("GetResourceCacheID", []interface{}{})
	fake.getResourceCacheIDMutex.Unlock()
	if fake.GetResourceCacheIDStub != nil {
		return fake.GetResourceCacheIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getResourceCacheIDReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) GetResourceCacheIDCallCount() int {
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	return len(fake.getResourceCacheIDArgsForCall)
}

func (fake *FakeVolume) GetResourceCacheIDCalls(stub func() int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = stub
}

func (fake *FakeVolume) GetResourceCacheIDReturns(result1 int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = nil
	fake.getResourceCacheIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeVolume) GetResourceCacheIDReturnsOnCall(i int, result1 int) {
	fake.getResourceCacheIDMutex.Lock()
	defer fake.getResourceCacheIDMutex.Unlock()
	fake.GetResourceCacheIDStub = nil
	if fake.getResourceCacheIDReturnsOnCall == nil {
		fake.getResourceCacheIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getResourceCacheIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) InitializeArtifact(arg1 string, arg2 int) (db.WorkerArtifact, error) {
	fake.initializeArtifactMutex.Lock()
	ret, specificReturn := fake.initializeArtifactReturnsOnCall[len(fake.initializeArtifactArgsForCall)]
	fake.initializeArtifactArgsForCall = append(fake.initializeArtifactArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("InitializeArtifact", []interface{}{arg1, arg2})
	fake.initializeArtifactMutex.Unlock()
	if fake.InitializeArtifactStub != nil {
		return fake.InitializeArtifactStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.initializeArtifactReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) InitializeArtifactCallCount() int {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	return len(fake.initializeArtifactArgsForCall)
}

func (fake *FakeVolume) InitializeArtifactCalls(stub func(string, int) (db.WorkerArtifact, error)) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = stub
}

func (fake *FakeVolume) InitializeArtifactArgsForCall(i int) (string, int) {
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	argsForCall := fake.initializeArtifactArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) InitializeArtifactReturns(result1 db.WorkerArtifact, result2 error) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = nil
	fake.initializeArtifactReturns = struct {
		result1 db.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) InitializeArtifactReturnsOnCall(i int, result1 db.WorkerArtifact, result2 error) {
	fake.initializeArtifactMutex.Lock()
	defer fake.initializeArtifactMutex.Unlock()
	fake.InitializeArtifactStub = nil
	if fake.initializeArtifactReturnsOnCall == nil {
		fake.initializeArtifactReturnsOnCall = make(map[int]struct {
			result1 db.WorkerArtifact
			result2 error
		})
	}
	fake.initializeArtifactReturnsOnCall[i] = struct {
		result1 db.WorkerArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) InitializeResourceCache(arg1 db.UsedResourceCache) error {
	fake.initializeResourceCacheMutex.Lock()
	ret, specificReturn := fake.initializeResourceCacheReturnsOnCall[len(fake.initializeResourceCacheArgsForCall)]
	fake.initializeResourceCacheArgsForCall = append(fake.initializeResourceCacheArgsForCall, struct {
		arg1 db.UsedResourceCache
	}{arg1})
	fake.recordInvocation("InitializeResourceCache", []interface{}{arg1})
	fake.initializeResourceCacheMutex.Unlock()
	if fake.InitializeResourceCacheStub != nil {
		return fake.InitializeResourceCacheStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializeResourceCacheReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) InitializeResourceCacheCallCount() int {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	return len(fake.initializeResourceCacheArgsForCall)
}

func (fake *FakeVolume) InitializeResourceCacheCalls(stub func(db.UsedResourceCache) error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = stub
}

func (fake *FakeVolume) InitializeResourceCacheArgsForCall(i int) db.UsedResourceCache {
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	argsForCall := fake.initializeResourceCacheArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) InitializeResourceCacheReturns(result1 error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = nil
	fake.initializeResourceCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeResourceCacheReturnsOnCall(i int, result1 error) {
	fake.initializeResourceCacheMutex.Lock()
	defer fake.initializeResourceCacheMutex.Unlock()
	fake.InitializeResourceCacheStub = nil
	if fake.initializeResourceCacheReturnsOnCall == nil {
		fake.initializeResourceCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeResourceCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeTaskCache(arg1 lager.Logger, arg2 int, arg3 string, arg4 string, arg5 bool) error {
	fake.initializeTaskCacheMutex.Lock()
	ret, specificReturn := fake.initializeTaskCacheReturnsOnCall[len(fake.initializeTaskCacheArgsForCall)]
	fake.initializeTaskCacheArgsForCall = append(fake.initializeTaskCacheArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("InitializeTaskCache", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.initializeTaskCacheMutex.Unlock()
	if fake.InitializeTaskCacheStub != nil {
		return fake.InitializeTaskCacheStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializeTaskCacheReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) InitializeTaskCacheCallCount() int {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	return len(fake.initializeTaskCacheArgsForCall)
}

func (fake *FakeVolume) InitializeTaskCacheCalls(stub func(lager.Logger, int, string, string, bool) error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = stub
}

func (fake *FakeVolume) InitializeTaskCacheArgsForCall(i int) (lager.Logger, int, string, string, bool) {
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	argsForCall := fake.initializeTaskCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVolume) InitializeTaskCacheReturns(result1 error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = nil
	fake.initializeTaskCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) InitializeTaskCacheReturnsOnCall(i int, result1 error) {
	fake.initializeTaskCacheMutex.Lock()
	defer fake.initializeTaskCacheMutex.Unlock()
	fake.InitializeTaskCacheStub = nil
	if fake.initializeTaskCacheReturnsOnCall == nil {
		fake.initializeTaskCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeTaskCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
	}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pathReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolume) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeVolume) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) PathReturnsOnCall(i int, result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Properties() (baggageclaim.VolumeProperties, error) {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
	}{})
	fake.recordInvocation("Properties", []interface{}{})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.propertiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeVolume) PropertiesCalls(stub func() (baggageclaim.VolumeProperties, error)) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = stub
}

func (fake *FakeVolume) PropertiesReturns(result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) PropertiesReturnsOnCall(i int, result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.VolumeProperties
			result2 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) SetPrivileged(arg1 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("SetPrivileged", []interface{}{arg1})
	fake.setPrivilegedMutex.Unlock()
	if fake.SetPrivilegedStub != nil {
		return fake.SetPrivilegedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPrivilegedReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeVolume) SetPrivilegedCalls(stub func(bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeVolume) SetPrivilegedArgsForCall(i int) bool {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetProperty(arg1 string, arg2 string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetProperty", []interface{}{arg1, arg2})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPropertyReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeVolume) SetPropertyCalls(stub func(string, string) error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = stub
}

func (fake *FakeVolume) SetPropertyArgsForCall(i int) (string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	argsForCall := fake.setPropertyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) SetPropertyReturns(result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamIn(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding, arg4 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamInReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVolume) StreamInCalls(stub func(context.Context, string, baggageclaim.Encoding, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeVolume) StreamInArgsForCall(i int) (context.Context, string, baggageclaim.Encoding, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamOut(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}{arg1, arg2, arg3})
	fake.recordInvocation("StreamOut", []interface{}{arg1, arg2, arg3})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.streamOutReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVolume) StreamOutCalls(stub func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeVolume) StreamOutArgsForCall(i int) (context.Context, string, baggageclaim.Encoding) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolume) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) WorkerName() string {
	fake.workerNameMutex.Lock()
	ret, specificReturn := fake.workerNameReturnsOnCall[len(fake.workerNameArgsForCall)]
	fake.workerNameArgsForCall = append(fake.workerNameArgsForCall, struct {
	}{})
	fake.recordInvocation("WorkerName", []interface{}{})
	fake.workerNameMutex.Unlock()
	if fake.WorkerNameStub != nil {
		return fake.WorkerNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.workerNameReturns
	return fakeReturns.result1
}

func (fake *FakeVolume) WorkerNameCallCount() int {
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	return len(fake.workerNameArgsForCall)
}

func (fake *FakeVolume) WorkerNameCalls(stub func() string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = stub
}

func (fake *FakeVolume) WorkerNameReturns(result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	fake.workerNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) WorkerNameReturnsOnCall(i int, result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	if fake.workerNameReturnsOnCall == nil {
		fake.workerNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.workerNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cOWStrategyMutex.RLock()
	defer fake.cOWStrategyMutex.RUnlock()
	fake.createChildForContainerMutex.RLock()
	defer fake.createChildForContainerMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.getResourceCacheIDMutex.RLock()
	defer fake.getResourceCacheIDMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.initializeArtifactMutex.RLock()
	defer fake.initializeArtifactMutex.RUnlock()
	fake.initializeResourceCacheMutex.RLock()
	defer fake.initializeResourceCacheMutex.RUnlock()
	fake.initializeTaskCacheMutex.RLock()
	defer fake.initializeTaskCacheMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolume) recordInvocation(key string, args []interface{}) {
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

var _ worker.Volume = new(FakeVolume)
