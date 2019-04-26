// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/brokerapi/domain"
)

type AutoFakeServiceBroker struct {
	BindStub        func(context.Context, string, string, domain.BindDetails, bool) (domain.Binding, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.BindDetails
		arg5 bool
	}
	bindReturns struct {
		result1 domain.Binding
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 domain.Binding
		result2 error
	}
	DeprovisionStub        func(context.Context, string, domain.DeprovisionDetails, bool) (domain.DeprovisionServiceSpec, error)
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.DeprovisionDetails
		arg4 bool
	}
	deprovisionReturns struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}
	GetBindingStub        func(context.Context, string, string) (domain.GetBindingSpec, error)
	getBindingMutex       sync.RWMutex
	getBindingArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	getBindingReturns struct {
		result1 domain.GetBindingSpec
		result2 error
	}
	getBindingReturnsOnCall map[int]struct {
		result1 domain.GetBindingSpec
		result2 error
	}
	GetInstanceStub        func(context.Context, string) (domain.GetInstanceDetailsSpec, error)
	getInstanceMutex       sync.RWMutex
	getInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getInstanceReturns struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}
	getInstanceReturnsOnCall map[int]struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}
	LastBindingOperationStub        func(context.Context, string, string, domain.PollDetails) (domain.LastOperation, error)
	lastBindingOperationMutex       sync.RWMutex
	lastBindingOperationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.PollDetails
	}
	lastBindingOperationReturns struct {
		result1 domain.LastOperation
		result2 error
	}
	lastBindingOperationReturnsOnCall map[int]struct {
		result1 domain.LastOperation
		result2 error
	}
	LastOperationStub        func(context.Context, string, domain.PollDetails) (domain.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.PollDetails
	}
	lastOperationReturns struct {
		result1 domain.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 domain.LastOperation
		result2 error
	}
	ProvisionStub        func(context.Context, string, domain.ProvisionDetails, bool) (domain.ProvisionedServiceSpec, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.ProvisionDetails
		arg4 bool
	}
	provisionReturns struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}
	ServicesStub        func(context.Context) ([]domain.Service, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
		arg1 context.Context
	}
	servicesReturns struct {
		result1 []domain.Service
		result2 error
	}
	servicesReturnsOnCall map[int]struct {
		result1 []domain.Service
		result2 error
	}
	UnbindStub        func(context.Context, string, string, domain.UnbindDetails, bool) (domain.UnbindSpec, error)
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.UnbindDetails
		arg5 bool
	}
	unbindReturns struct {
		result1 domain.UnbindSpec
		result2 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 domain.UnbindSpec
		result2 error
	}
	UpdateStub        func(context.Context, string, domain.UpdateDetails, bool) (domain.UpdateServiceSpec, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 domain.UpdateDetails
		arg4 bool
	}
	updateReturns struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AutoFakeServiceBroker) Bind(arg1 context.Context, arg2 string, arg3 string, arg4 domain.BindDetails, arg5 bool) (domain.Binding, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.BindDetails
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Bind", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.bindReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *AutoFakeServiceBroker) BindCalls(stub func(context.Context, string, string, domain.BindDetails, bool) (domain.Binding, error)) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = stub
}

func (fake *AutoFakeServiceBroker) BindArgsForCall(i int) (context.Context, string, string, domain.BindDetails, bool) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	argsForCall := fake.bindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *AutoFakeServiceBroker) BindReturns(result1 domain.Binding, result2 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 domain.Binding
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) BindReturnsOnCall(i int, result1 domain.Binding, result2 error) {
	fake.bindMutex.Lock()
	defer fake.bindMutex.Unlock()
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 domain.Binding
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 domain.Binding
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Deprovision(arg1 context.Context, arg2 string, arg3 domain.DeprovisionDetails, arg4 bool) (domain.DeprovisionServiceSpec, error) {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.DeprovisionDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Deprovision", []interface{}{arg1, arg2, arg3, arg4})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deprovisionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *AutoFakeServiceBroker) DeprovisionCalls(stub func(context.Context, string, domain.DeprovisionDetails, bool) (domain.DeprovisionServiceSpec, error)) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = stub
}

func (fake *AutoFakeServiceBroker) DeprovisionArgsForCall(i int) (context.Context, string, domain.DeprovisionDetails, bool) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	argsForCall := fake.deprovisionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *AutoFakeServiceBroker) DeprovisionReturns(result1 domain.DeprovisionServiceSpec, result2 error) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) DeprovisionReturnsOnCall(i int, result1 domain.DeprovisionServiceSpec, result2 error) {
	fake.deprovisionMutex.Lock()
	defer fake.deprovisionMutex.Unlock()
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 domain.DeprovisionServiceSpec
			result2 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 domain.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) GetBinding(arg1 context.Context, arg2 string, arg3 string) (domain.GetBindingSpec, error) {
	fake.getBindingMutex.Lock()
	ret, specificReturn := fake.getBindingReturnsOnCall[len(fake.getBindingArgsForCall)]
	fake.getBindingArgsForCall = append(fake.getBindingArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetBinding", []interface{}{arg1, arg2, arg3})
	fake.getBindingMutex.Unlock()
	if fake.GetBindingStub != nil {
		return fake.GetBindingStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBindingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) GetBindingCallCount() int {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return len(fake.getBindingArgsForCall)
}

func (fake *AutoFakeServiceBroker) GetBindingCalls(stub func(context.Context, string, string) (domain.GetBindingSpec, error)) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = stub
}

func (fake *AutoFakeServiceBroker) GetBindingArgsForCall(i int) (context.Context, string, string) {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	argsForCall := fake.getBindingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *AutoFakeServiceBroker) GetBindingReturns(result1 domain.GetBindingSpec, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	fake.getBindingReturns = struct {
		result1 domain.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) GetBindingReturnsOnCall(i int, result1 domain.GetBindingSpec, result2 error) {
	fake.getBindingMutex.Lock()
	defer fake.getBindingMutex.Unlock()
	fake.GetBindingStub = nil
	if fake.getBindingReturnsOnCall == nil {
		fake.getBindingReturnsOnCall = make(map[int]struct {
			result1 domain.GetBindingSpec
			result2 error
		})
	}
	fake.getBindingReturnsOnCall[i] = struct {
		result1 domain.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) GetInstance(arg1 context.Context, arg2 string) (domain.GetInstanceDetailsSpec, error) {
	fake.getInstanceMutex.Lock()
	ret, specificReturn := fake.getInstanceReturnsOnCall[len(fake.getInstanceArgsForCall)]
	fake.getInstanceArgsForCall = append(fake.getInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetInstance", []interface{}{arg1, arg2})
	fake.getInstanceMutex.Unlock()
	if fake.GetInstanceStub != nil {
		return fake.GetInstanceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstanceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) GetInstanceCallCount() int {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	return len(fake.getInstanceArgsForCall)
}

func (fake *AutoFakeServiceBroker) GetInstanceCalls(stub func(context.Context, string) (domain.GetInstanceDetailsSpec, error)) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = stub
}

func (fake *AutoFakeServiceBroker) GetInstanceArgsForCall(i int) (context.Context, string) {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	argsForCall := fake.getInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *AutoFakeServiceBroker) GetInstanceReturns(result1 domain.GetInstanceDetailsSpec, result2 error) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = nil
	fake.getInstanceReturns = struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) GetInstanceReturnsOnCall(i int, result1 domain.GetInstanceDetailsSpec, result2 error) {
	fake.getInstanceMutex.Lock()
	defer fake.getInstanceMutex.Unlock()
	fake.GetInstanceStub = nil
	if fake.getInstanceReturnsOnCall == nil {
		fake.getInstanceReturnsOnCall = make(map[int]struct {
			result1 domain.GetInstanceDetailsSpec
			result2 error
		})
	}
	fake.getInstanceReturnsOnCall[i] = struct {
		result1 domain.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) LastBindingOperation(arg1 context.Context, arg2 string, arg3 string, arg4 domain.PollDetails) (domain.LastOperation, error) {
	fake.lastBindingOperationMutex.Lock()
	ret, specificReturn := fake.lastBindingOperationReturnsOnCall[len(fake.lastBindingOperationArgsForCall)]
	fake.lastBindingOperationArgsForCall = append(fake.lastBindingOperationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.PollDetails
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("LastBindingOperation", []interface{}{arg1, arg2, arg3, arg4})
	fake.lastBindingOperationMutex.Unlock()
	if fake.LastBindingOperationStub != nil {
		return fake.LastBindingOperationStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lastBindingOperationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) LastBindingOperationCallCount() int {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	return len(fake.lastBindingOperationArgsForCall)
}

func (fake *AutoFakeServiceBroker) LastBindingOperationCalls(stub func(context.Context, string, string, domain.PollDetails) (domain.LastOperation, error)) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = stub
}

func (fake *AutoFakeServiceBroker) LastBindingOperationArgsForCall(i int) (context.Context, string, string, domain.PollDetails) {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	argsForCall := fake.lastBindingOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *AutoFakeServiceBroker) LastBindingOperationReturns(result1 domain.LastOperation, result2 error) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = nil
	fake.lastBindingOperationReturns = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) LastBindingOperationReturnsOnCall(i int, result1 domain.LastOperation, result2 error) {
	fake.lastBindingOperationMutex.Lock()
	defer fake.lastBindingOperationMutex.Unlock()
	fake.LastBindingOperationStub = nil
	if fake.lastBindingOperationReturnsOnCall == nil {
		fake.lastBindingOperationReturnsOnCall = make(map[int]struct {
			result1 domain.LastOperation
			result2 error
		})
	}
	fake.lastBindingOperationReturnsOnCall[i] = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) LastOperation(arg1 context.Context, arg2 string, arg3 domain.PollDetails) (domain.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.PollDetails
	}{arg1, arg2, arg3})
	fake.recordInvocation("LastOperation", []interface{}{arg1, arg2, arg3})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lastOperationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *AutoFakeServiceBroker) LastOperationCalls(stub func(context.Context, string, domain.PollDetails) (domain.LastOperation, error)) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = stub
}

func (fake *AutoFakeServiceBroker) LastOperationArgsForCall(i int) (context.Context, string, domain.PollDetails) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	argsForCall := fake.lastOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *AutoFakeServiceBroker) LastOperationReturns(result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) LastOperationReturnsOnCall(i int, result1 domain.LastOperation, result2 error) {
	fake.lastOperationMutex.Lock()
	defer fake.lastOperationMutex.Unlock()
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 domain.LastOperation
			result2 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 domain.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Provision(arg1 context.Context, arg2 string, arg3 domain.ProvisionDetails, arg4 bool) (domain.ProvisionedServiceSpec, error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.ProvisionDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Provision", []interface{}{arg1, arg2, arg3, arg4})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.provisionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *AutoFakeServiceBroker) ProvisionCalls(stub func(context.Context, string, domain.ProvisionDetails, bool) (domain.ProvisionedServiceSpec, error)) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = stub
}

func (fake *AutoFakeServiceBroker) ProvisionArgsForCall(i int) (context.Context, string, domain.ProvisionDetails, bool) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	argsForCall := fake.provisionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *AutoFakeServiceBroker) ProvisionReturns(result1 domain.ProvisionedServiceSpec, result2 error) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) ProvisionReturnsOnCall(i int, result1 domain.ProvisionedServiceSpec, result2 error) {
	fake.provisionMutex.Lock()
	defer fake.provisionMutex.Unlock()
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 domain.ProvisionedServiceSpec
			result2 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 domain.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Services(arg1 context.Context) ([]domain.Service, error) {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Services", []interface{}{arg1})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.servicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *AutoFakeServiceBroker) ServicesCalls(stub func(context.Context) ([]domain.Service, error)) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = stub
}

func (fake *AutoFakeServiceBroker) ServicesArgsForCall(i int) context.Context {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	argsForCall := fake.servicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AutoFakeServiceBroker) ServicesReturns(result1 []domain.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []domain.Service
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) ServicesReturnsOnCall(i int, result1 []domain.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []domain.Service
			result2 error
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []domain.Service
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Unbind(arg1 context.Context, arg2 string, arg3 string, arg4 domain.UnbindDetails, arg5 bool) (domain.UnbindSpec, error) {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 domain.UnbindDetails
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Unbind", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unbindReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *AutoFakeServiceBroker) UnbindCalls(stub func(context.Context, string, string, domain.UnbindDetails, bool) (domain.UnbindSpec, error)) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = stub
}

func (fake *AutoFakeServiceBroker) UnbindArgsForCall(i int) (context.Context, string, string, domain.UnbindDetails, bool) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	argsForCall := fake.unbindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *AutoFakeServiceBroker) UnbindReturns(result1 domain.UnbindSpec, result2 error) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 domain.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) UnbindReturnsOnCall(i int, result1 domain.UnbindSpec, result2 error) {
	fake.unbindMutex.Lock()
	defer fake.unbindMutex.Unlock()
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 domain.UnbindSpec
			result2 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 domain.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Update(arg1 context.Context, arg2 string, arg3 domain.UpdateDetails, arg4 bool) (domain.UpdateServiceSpec, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 domain.UpdateDetails
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AutoFakeServiceBroker) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *AutoFakeServiceBroker) UpdateCalls(stub func(context.Context, string, domain.UpdateDetails, bool) (domain.UpdateServiceSpec, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *AutoFakeServiceBroker) UpdateArgsForCall(i int) (context.Context, string, domain.UpdateDetails, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *AutoFakeServiceBroker) UpdateReturns(result1 domain.UpdateServiceSpec, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) UpdateReturnsOnCall(i int, result1 domain.UpdateServiceSpec, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 domain.UpdateServiceSpec
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 domain.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *AutoFakeServiceBroker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AutoFakeServiceBroker) recordInvocation(key string, args []interface{}) {
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

var _ brokerapi.ServiceBroker = new(AutoFakeServiceBroker)
