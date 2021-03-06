// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/chenbh/concourse/atc/db"
)

type FakePipelineLifecycle struct {
	ArchiveAbandonedPipelinesStub        func() error
	archiveAbandonedPipelinesMutex       sync.RWMutex
	archiveAbandonedPipelinesArgsForCall []struct {
	}
	archiveAbandonedPipelinesReturns struct {
		result1 error
	}
	archiveAbandonedPipelinesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelineLifecycle) ArchiveAbandonedPipelines() error {
	fake.archiveAbandonedPipelinesMutex.Lock()
	ret, specificReturn := fake.archiveAbandonedPipelinesReturnsOnCall[len(fake.archiveAbandonedPipelinesArgsForCall)]
	fake.archiveAbandonedPipelinesArgsForCall = append(fake.archiveAbandonedPipelinesArgsForCall, struct {
	}{})
	fake.recordInvocation("ArchiveAbandonedPipelines", []interface{}{})
	fake.archiveAbandonedPipelinesMutex.Unlock()
	if fake.ArchiveAbandonedPipelinesStub != nil {
		return fake.ArchiveAbandonedPipelinesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.archiveAbandonedPipelinesReturns
	return fakeReturns.result1
}

func (fake *FakePipelineLifecycle) ArchiveAbandonedPipelinesCallCount() int {
	fake.archiveAbandonedPipelinesMutex.RLock()
	defer fake.archiveAbandonedPipelinesMutex.RUnlock()
	return len(fake.archiveAbandonedPipelinesArgsForCall)
}

func (fake *FakePipelineLifecycle) ArchiveAbandonedPipelinesCalls(stub func() error) {
	fake.archiveAbandonedPipelinesMutex.Lock()
	defer fake.archiveAbandonedPipelinesMutex.Unlock()
	fake.ArchiveAbandonedPipelinesStub = stub
}

func (fake *FakePipelineLifecycle) ArchiveAbandonedPipelinesReturns(result1 error) {
	fake.archiveAbandonedPipelinesMutex.Lock()
	defer fake.archiveAbandonedPipelinesMutex.Unlock()
	fake.ArchiveAbandonedPipelinesStub = nil
	fake.archiveAbandonedPipelinesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePipelineLifecycle) ArchiveAbandonedPipelinesReturnsOnCall(i int, result1 error) {
	fake.archiveAbandonedPipelinesMutex.Lock()
	defer fake.archiveAbandonedPipelinesMutex.Unlock()
	fake.ArchiveAbandonedPipelinesStub = nil
	if fake.archiveAbandonedPipelinesReturnsOnCall == nil {
		fake.archiveAbandonedPipelinesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.archiveAbandonedPipelinesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePipelineLifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.archiveAbandonedPipelinesMutex.RLock()
	defer fake.archiveAbandonedPipelinesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePipelineLifecycle) recordInvocation(key string, args []interface{}) {
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

var _ db.PipelineLifecycle = new(FakePipelineLifecycle)
