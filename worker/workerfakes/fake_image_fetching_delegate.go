// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"io"
	"sync"

	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"
	"github.com/chenbh/concourse/atc/worker"
)

type FakeImageFetchingDelegate struct {
	ImageVersionDeterminedStub        func(db.UsedResourceCache) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 db.UsedResourceCache
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	imageVersionDeterminedReturnsOnCall map[int]struct {
		result1 error
	}
	RedactImageSourceStub        func(atc.Source) (atc.Source, error)
	redactImageSourceMutex       sync.RWMutex
	redactImageSourceArgsForCall []struct {
		arg1 atc.Source
	}
	redactImageSourceReturns struct {
		result1 atc.Source
		result2 error
	}
	redactImageSourceReturnsOnCall map[int]struct {
		result1 atc.Source
		result2 error
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFetchingDelegate) ImageVersionDetermined(arg1 db.UsedResourceCache) error {
	fake.imageVersionDeterminedMutex.Lock()
	ret, specificReturn := fake.imageVersionDeterminedReturnsOnCall[len(fake.imageVersionDeterminedArgsForCall)]
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 db.UsedResourceCache
	}{arg1})
	fake.recordInvocation("ImageVersionDetermined", []interface{}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.imageVersionDeterminedReturns
	return fakeReturns.result1
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedCalls(stub func(db.UsedResourceCache) error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = stub
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedArgsForCall(i int) db.UsedResourceCache {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	argsForCall := fake.imageVersionDeterminedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedReturnsOnCall(i int, result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	if fake.imageVersionDeterminedReturnsOnCall == nil {
		fake.imageVersionDeterminedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.imageVersionDeterminedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFetchingDelegate) RedactImageSource(arg1 atc.Source) (atc.Source, error) {
	fake.redactImageSourceMutex.Lock()
	ret, specificReturn := fake.redactImageSourceReturnsOnCall[len(fake.redactImageSourceArgsForCall)]
	fake.redactImageSourceArgsForCall = append(fake.redactImageSourceArgsForCall, struct {
		arg1 atc.Source
	}{arg1})
	fake.recordInvocation("RedactImageSource", []interface{}{arg1})
	fake.redactImageSourceMutex.Unlock()
	if fake.RedactImageSourceStub != nil {
		return fake.RedactImageSourceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.redactImageSourceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageFetchingDelegate) RedactImageSourceCallCount() int {
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	return len(fake.redactImageSourceArgsForCall)
}

func (fake *FakeImageFetchingDelegate) RedactImageSourceCalls(stub func(atc.Source) (atc.Source, error)) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = stub
}

func (fake *FakeImageFetchingDelegate) RedactImageSourceArgsForCall(i int) atc.Source {
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	argsForCall := fake.redactImageSourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImageFetchingDelegate) RedactImageSourceReturns(result1 atc.Source, result2 error) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = nil
	fake.redactImageSourceReturns = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFetchingDelegate) RedactImageSourceReturnsOnCall(i int, result1 atc.Source, result2 error) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = nil
	if fake.redactImageSourceReturnsOnCall == nil {
		fake.redactImageSourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
			result2 error
		})
	}
	fake.redactImageSourceReturnsOnCall[i] = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFetchingDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stderrReturns
	return fakeReturns.result1
}

func (fake *FakeImageFetchingDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeImageFetchingDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeImageFetchingDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stdoutReturns
	return fakeReturns.result1
}

func (fake *FakeImageFetchingDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeImageFetchingDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeImageFetchingDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageFetchingDelegate) recordInvocation(key string, args []interface{}) {
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

var _ worker.ImageFetchingDelegate = new(FakeImageFetchingDelegate)
