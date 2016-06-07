// This file was generated by counterfeiter
package fake_service_adapter

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
)

type FakeBinder struct {
	CreateBindingStub        func(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, arbitraryParams map[string]interface{}) (serviceadapter.Binding, error)
	createBindingMutex       sync.RWMutex
	createBindingArgsForCall []struct {
		bindingID          string
		deploymentTopology bosh.BoshVMs
		manifest           bosh.BoshManifest
		arbitraryParams    map[string]interface{}
	}
	createBindingReturns struct {
		result1 serviceadapter.Binding
		result2 error
	}
	DeleteBindingStub        func(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) error
	deleteBindingMutex       sync.RWMutex
	deleteBindingArgsForCall []struct {
		bindingID          string
		deploymentTopology bosh.BoshVMs
		manifest           bosh.BoshManifest
	}
	deleteBindingReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBinder) Create(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, arbitraryParams map[string]interface{}) (serviceadapter.Binding, error) {
	fake.createBindingMutex.Lock()
	fake.createBindingArgsForCall = append(fake.createBindingArgsForCall, struct {
		bindingID          string
		deploymentTopology bosh.BoshVMs
		manifest           bosh.BoshManifest
		arbitraryParams    map[string]interface{}
	}{bindingID, deploymentTopology, manifest, arbitraryParams})
	fake.recordInvocation("CreateBinding", []interface{}{bindingID, deploymentTopology, manifest, arbitraryParams})
	fake.createBindingMutex.Unlock()
	if fake.CreateBindingStub != nil {
		return fake.CreateBindingStub(bindingID, deploymentTopology, manifest, arbitraryParams)
	} else {
		return fake.createBindingReturns.result1, fake.createBindingReturns.result2
	}
}

func (fake *FakeBinder) CreateBindingCallCount() int {
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	return len(fake.createBindingArgsForCall)
}

func (fake *FakeBinder) CreateBindingArgsForCall(i int) (string, bosh.BoshVMs, bosh.BoshManifest, map[string]interface{}) {
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	return fake.createBindingArgsForCall[i].bindingID, fake.createBindingArgsForCall[i].deploymentTopology, fake.createBindingArgsForCall[i].manifest, fake.createBindingArgsForCall[i].arbitraryParams
}

func (fake *FakeBinder) CreateBindingReturns(result1 serviceadapter.Binding, result2 error) {
	fake.CreateBindingStub = nil
	fake.createBindingReturns = struct {
		result1 serviceadapter.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeBinder) Delete(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) error {
	fake.deleteBindingMutex.Lock()
	fake.deleteBindingArgsForCall = append(fake.deleteBindingArgsForCall, struct {
		bindingID          string
		deploymentTopology bosh.BoshVMs
		manifest           bosh.BoshManifest
	}{bindingID, deploymentTopology, manifest})
	fake.recordInvocation("DeleteBinding", []interface{}{bindingID, deploymentTopology, manifest})
	fake.deleteBindingMutex.Unlock()
	if fake.DeleteBindingStub != nil {
		return fake.DeleteBindingStub(bindingID, deploymentTopology, manifest)
	} else {
		return fake.deleteBindingReturns.result1
	}
}

func (fake *FakeBinder) DeleteBindingCallCount() int {
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return len(fake.deleteBindingArgsForCall)
}

func (fake *FakeBinder) DeleteBindingArgsForCall(i int) (string, bosh.BoshVMs, bosh.BoshManifest) {
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return fake.deleteBindingArgsForCall[i].bindingID, fake.deleteBindingArgsForCall[i].deploymentTopology, fake.deleteBindingArgsForCall[i].manifest
}

func (fake *FakeBinder) DeleteBindingReturns(result1 error) {
	fake.DeleteBindingStub = nil
	fake.deleteBindingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBinder) recordInvocation(key string, args []interface{}) {
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

var _ serviceadapter.Binder = new(FakeBinder)
