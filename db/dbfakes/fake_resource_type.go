// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

type FakeResourceType struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct{}
	typeReturns     struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	PrivilegedStub        func() bool
	privilegedMutex       sync.RWMutex
	privilegedArgsForCall []struct{}
	privilegedReturns     struct {
		result1 bool
	}
	privilegedReturnsOnCall map[int]struct {
		result1 bool
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct{}
	sourceReturns     struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	SaveVersionStub        func(atc.Version) error
	saveVersionMutex       sync.RWMutex
	saveVersionArgsForCall []struct {
		arg1 atc.Version
	}
	saveVersionReturns struct {
		result1 error
	}
	saveVersionReturnsOnCall map[int]struct {
		result1 error
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceType) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeResourceType) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceType) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceType) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceType) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeResourceType) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeResourceType) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct{}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.typeReturns.result1
}

func (fake *FakeResourceType) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeResourceType) TypeReturns(result1 string) {
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) TypeReturnsOnCall(i int, result1 string) {
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourceType) Privileged() bool {
	fake.privilegedMutex.Lock()
	ret, specificReturn := fake.privilegedReturnsOnCall[len(fake.privilegedArgsForCall)]
	fake.privilegedArgsForCall = append(fake.privilegedArgsForCall, struct{}{})
	fake.recordInvocation("Privileged", []interface{}{})
	fake.privilegedMutex.Unlock()
	if fake.PrivilegedStub != nil {
		return fake.PrivilegedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.privilegedReturns.result1
}

func (fake *FakeResourceType) PrivilegedCallCount() int {
	fake.privilegedMutex.RLock()
	defer fake.privilegedMutex.RUnlock()
	return len(fake.privilegedArgsForCall)
}

func (fake *FakeResourceType) PrivilegedReturns(result1 bool) {
	fake.PrivilegedStub = nil
	fake.privilegedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResourceType) PrivilegedReturnsOnCall(i int, result1 bool) {
	fake.PrivilegedStub = nil
	if fake.privilegedReturnsOnCall == nil {
		fake.privilegedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.privilegedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeResourceType) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct{}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sourceReturns.result1
}

func (fake *FakeResourceType) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeResourceType) SourceReturns(result1 atc.Source) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceType) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceType) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeResourceType) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceType) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceType) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceType) SaveVersion(arg1 atc.Version) error {
	fake.saveVersionMutex.Lock()
	ret, specificReturn := fake.saveVersionReturnsOnCall[len(fake.saveVersionArgsForCall)]
	fake.saveVersionArgsForCall = append(fake.saveVersionArgsForCall, struct {
		arg1 atc.Version
	}{arg1})
	fake.recordInvocation("SaveVersion", []interface{}{arg1})
	fake.saveVersionMutex.Unlock()
	if fake.SaveVersionStub != nil {
		return fake.SaveVersionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveVersionReturns.result1
}

func (fake *FakeResourceType) SaveVersionCallCount() int {
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	return len(fake.saveVersionArgsForCall)
}

func (fake *FakeResourceType) SaveVersionArgsForCall(i int) atc.Version {
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	return fake.saveVersionArgsForCall[i].arg1
}

func (fake *FakeResourceType) SaveVersionReturns(result1 error) {
	fake.SaveVersionStub = nil
	fake.saveVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) SaveVersionReturnsOnCall(i int, result1 error) {
	fake.SaveVersionStub = nil
	if fake.saveVersionReturnsOnCall == nil {
		fake.saveVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceType) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.reloadReturns.result1, fake.reloadReturns.result2
}

func (fake *FakeResourceType) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeResourceType) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceType) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.privilegedMutex.RLock()
	defer fake.privilegedMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceType) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceType = new(FakeResourceType)
