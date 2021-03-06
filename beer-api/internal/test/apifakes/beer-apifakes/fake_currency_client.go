// Code generated by counterfeiter. DO NOT EDIT.
package beerapifakes

import (
	"context"
	"sync"

	v1 "github.com/jkarlos000/technical-challenge/currency/api/proto/v1"
	"google.golang.org/grpc"
)

type FakeCurrencyClient struct {
	GetPriceStub        func(context.Context, *v1.Request, ...grpc.CallOption) (*v1.Response, error)
	getPriceMutex       sync.RWMutex
	getPriceArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Request
		arg3 []grpc.CallOption
	}
	getPriceReturns struct {
		result1 *v1.Response
		result2 error
	}
	getPriceReturnsOnCall map[int]struct {
		result1 *v1.Response
		result2 error
	}
	GetPriceStreamStub        func(context.Context, *v1.Request, ...grpc.CallOption) (v1.Currency_GetPriceStreamClient, error)
	getPriceStreamMutex       sync.RWMutex
	getPriceStreamArgsForCall []struct {
		arg1 context.Context
		arg2 *v1.Request
		arg3 []grpc.CallOption
	}
	getPriceStreamReturns struct {
		result1 v1.Currency_GetPriceStreamClient
		result2 error
	}
	getPriceStreamReturnsOnCall map[int]struct {
		result1 v1.Currency_GetPriceStreamClient
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCurrencyClient) GetPrice(arg1 context.Context, arg2 *v1.Request, arg3 ...grpc.CallOption) (*v1.Response, error) {
	fake.getPriceMutex.Lock()
	ret, specificReturn := fake.getPriceReturnsOnCall[len(fake.getPriceArgsForCall)]
	fake.getPriceArgsForCall = append(fake.getPriceArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Request
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GetPriceStub
	fakeReturns := fake.getPriceReturns
	fake.recordInvocation("GetPrice", []interface{}{arg1, arg2, arg3})
	fake.getPriceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCurrencyClient) GetPriceCallCount() int {
	fake.getPriceMutex.RLock()
	defer fake.getPriceMutex.RUnlock()
	return len(fake.getPriceArgsForCall)
}

func (fake *FakeCurrencyClient) GetPriceCalls(stub func(context.Context, *v1.Request, ...grpc.CallOption) (*v1.Response, error)) {
	fake.getPriceMutex.Lock()
	defer fake.getPriceMutex.Unlock()
	fake.GetPriceStub = stub
}

func (fake *FakeCurrencyClient) GetPriceArgsForCall(i int) (context.Context, *v1.Request, []grpc.CallOption) {
	fake.getPriceMutex.RLock()
	defer fake.getPriceMutex.RUnlock()
	argsForCall := fake.getPriceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCurrencyClient) GetPriceReturns(result1 *v1.Response, result2 error) {
	fake.getPriceMutex.Lock()
	defer fake.getPriceMutex.Unlock()
	fake.GetPriceStub = nil
	fake.getPriceReturns = struct {
		result1 *v1.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeCurrencyClient) GetPriceReturnsOnCall(i int, result1 *v1.Response, result2 error) {
	fake.getPriceMutex.Lock()
	defer fake.getPriceMutex.Unlock()
	fake.GetPriceStub = nil
	if fake.getPriceReturnsOnCall == nil {
		fake.getPriceReturnsOnCall = make(map[int]struct {
			result1 *v1.Response
			result2 error
		})
	}
	fake.getPriceReturnsOnCall[i] = struct {
		result1 *v1.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeCurrencyClient) GetPriceStream(arg1 context.Context, arg2 *v1.Request, arg3 ...grpc.CallOption) (v1.Currency_GetPriceStreamClient, error) {
	fake.getPriceStreamMutex.Lock()
	ret, specificReturn := fake.getPriceStreamReturnsOnCall[len(fake.getPriceStreamArgsForCall)]
	fake.getPriceStreamArgsForCall = append(fake.getPriceStreamArgsForCall, struct {
		arg1 context.Context
		arg2 *v1.Request
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GetPriceStreamStub
	fakeReturns := fake.getPriceStreamReturns
	fake.recordInvocation("GetPriceStream", []interface{}{arg1, arg2, arg3})
	fake.getPriceStreamMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCurrencyClient) GetPriceStreamCallCount() int {
	fake.getPriceStreamMutex.RLock()
	defer fake.getPriceStreamMutex.RUnlock()
	return len(fake.getPriceStreamArgsForCall)
}

func (fake *FakeCurrencyClient) GetPriceStreamCalls(stub func(context.Context, *v1.Request, ...grpc.CallOption) (v1.Currency_GetPriceStreamClient, error)) {
	fake.getPriceStreamMutex.Lock()
	defer fake.getPriceStreamMutex.Unlock()
	fake.GetPriceStreamStub = stub
}

func (fake *FakeCurrencyClient) GetPriceStreamArgsForCall(i int) (context.Context, *v1.Request, []grpc.CallOption) {
	fake.getPriceStreamMutex.RLock()
	defer fake.getPriceStreamMutex.RUnlock()
	argsForCall := fake.getPriceStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCurrencyClient) GetPriceStreamReturns(result1 v1.Currency_GetPriceStreamClient, result2 error) {
	fake.getPriceStreamMutex.Lock()
	defer fake.getPriceStreamMutex.Unlock()
	fake.GetPriceStreamStub = nil
	fake.getPriceStreamReturns = struct {
		result1 v1.Currency_GetPriceStreamClient
		result2 error
	}{result1, result2}
}

func (fake *FakeCurrencyClient) GetPriceStreamReturnsOnCall(i int, result1 v1.Currency_GetPriceStreamClient, result2 error) {
	fake.getPriceStreamMutex.Lock()
	defer fake.getPriceStreamMutex.Unlock()
	fake.GetPriceStreamStub = nil
	if fake.getPriceStreamReturnsOnCall == nil {
		fake.getPriceStreamReturnsOnCall = make(map[int]struct {
			result1 v1.Currency_GetPriceStreamClient
			result2 error
		})
	}
	fake.getPriceStreamReturnsOnCall[i] = struct {
		result1 v1.Currency_GetPriceStreamClient
		result2 error
	}{result1, result2}
}

func (fake *FakeCurrencyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPriceMutex.RLock()
	defer fake.getPriceMutex.RUnlock()
	fake.getPriceStreamMutex.RLock()
	defer fake.getPriceStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCurrencyClient) recordInvocation(key string, args []interface{}) {
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

var _ v1.CurrencyClient = new(FakeCurrencyClient)
