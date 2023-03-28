package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain.ProductServiceClient -o ./mocks/product_service_client_minimock.go -n ProductServiceClientMock

import (
	"context"
	mm_domain "route256/checkout/internal/domain"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ProductServiceClientMock implements domain.ProductServiceClient
type ProductServiceClientMock struct {
	t minimock.Tester

	funcProduct          func(ctx context.Context, sku uint32) (p1 mm_domain.Product, err error)
	inspectFuncProduct   func(ctx context.Context, sku uint32)
	afterProductCounter  uint64
	beforeProductCounter uint64
	ProductMock          mProductServiceClientMockProduct
}

// NewProductServiceClientMock returns a mock for domain.ProductServiceClient
func NewProductServiceClientMock(t minimock.Tester) *ProductServiceClientMock {
	m := &ProductServiceClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ProductMock = mProductServiceClientMockProduct{mock: m}
	m.ProductMock.callArgs = []*ProductServiceClientMockProductParams{}

	return m
}

type mProductServiceClientMockProduct struct {
	mock               *ProductServiceClientMock
	defaultExpectation *ProductServiceClientMockProductExpectation
	expectations       []*ProductServiceClientMockProductExpectation

	callArgs []*ProductServiceClientMockProductParams
	mutex    sync.RWMutex
}

// ProductServiceClientMockProductExpectation specifies expectation struct of the ProductServiceClient.Product
type ProductServiceClientMockProductExpectation struct {
	mock    *ProductServiceClientMock
	params  *ProductServiceClientMockProductParams
	results *ProductServiceClientMockProductResults
	Counter uint64
}

// ProductServiceClientMockProductParams contains parameters of the ProductServiceClient.Product
type ProductServiceClientMockProductParams struct {
	ctx context.Context
	sku uint32
}

// ProductServiceClientMockProductResults contains results of the ProductServiceClient.Product
type ProductServiceClientMockProductResults struct {
	p1  mm_domain.Product
	err error
}

// Expect sets up expected params for ProductServiceClient.Product
func (mmProduct *mProductServiceClientMockProduct) Expect(ctx context.Context, sku uint32) *mProductServiceClientMockProduct {
	if mmProduct.mock.funcProduct != nil {
		mmProduct.mock.t.Fatalf("ProductServiceClientMock.Product mock is already set by Set")
	}

	if mmProduct.defaultExpectation == nil {
		mmProduct.defaultExpectation = &ProductServiceClientMockProductExpectation{}
	}

	mmProduct.defaultExpectation.params = &ProductServiceClientMockProductParams{ctx, sku}
	for _, e := range mmProduct.expectations {
		if minimock.Equal(e.params, mmProduct.defaultExpectation.params) {
			mmProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmProduct.defaultExpectation.params)
		}
	}

	return mmProduct
}

// Inspect accepts an inspector function that has same arguments as the ProductServiceClient.Product
func (mmProduct *mProductServiceClientMockProduct) Inspect(f func(ctx context.Context, sku uint32)) *mProductServiceClientMockProduct {
	if mmProduct.mock.inspectFuncProduct != nil {
		mmProduct.mock.t.Fatalf("Inspect function is already set for ProductServiceClientMock.Product")
	}

	mmProduct.mock.inspectFuncProduct = f

	return mmProduct
}

// Return sets up results that will be returned by ProductServiceClient.Product
func (mmProduct *mProductServiceClientMockProduct) Return(p1 mm_domain.Product, err error) *ProductServiceClientMock {
	if mmProduct.mock.funcProduct != nil {
		mmProduct.mock.t.Fatalf("ProductServiceClientMock.Product mock is already set by Set")
	}

	if mmProduct.defaultExpectation == nil {
		mmProduct.defaultExpectation = &ProductServiceClientMockProductExpectation{mock: mmProduct.mock}
	}
	mmProduct.defaultExpectation.results = &ProductServiceClientMockProductResults{p1, err}
	return mmProduct.mock
}

// Set uses given function f to mock the ProductServiceClient.Product method
func (mmProduct *mProductServiceClientMockProduct) Set(f func(ctx context.Context, sku uint32) (p1 mm_domain.Product, err error)) *ProductServiceClientMock {
	if mmProduct.defaultExpectation != nil {
		mmProduct.mock.t.Fatalf("Default expectation is already set for the ProductServiceClient.Product method")
	}

	if len(mmProduct.expectations) > 0 {
		mmProduct.mock.t.Fatalf("Some expectations are already set for the ProductServiceClient.Product method")
	}

	mmProduct.mock.funcProduct = f
	return mmProduct.mock
}

// When sets expectation for the ProductServiceClient.Product which will trigger the result defined by the following
// Then helper
func (mmProduct *mProductServiceClientMockProduct) When(ctx context.Context, sku uint32) *ProductServiceClientMockProductExpectation {
	if mmProduct.mock.funcProduct != nil {
		mmProduct.mock.t.Fatalf("ProductServiceClientMock.Product mock is already set by Set")
	}

	expectation := &ProductServiceClientMockProductExpectation{
		mock:   mmProduct.mock,
		params: &ProductServiceClientMockProductParams{ctx, sku},
	}
	mmProduct.expectations = append(mmProduct.expectations, expectation)
	return expectation
}

// Then sets up ProductServiceClient.Product return parameters for the expectation previously defined by the When method
func (e *ProductServiceClientMockProductExpectation) Then(p1 mm_domain.Product, err error) *ProductServiceClientMock {
	e.results = &ProductServiceClientMockProductResults{p1, err}
	return e.mock
}

// Product implements domain.ProductServiceClient
func (mmProduct *ProductServiceClientMock) Product(ctx context.Context, sku uint32) (p1 mm_domain.Product, err error) {
	mm_atomic.AddUint64(&mmProduct.beforeProductCounter, 1)
	defer mm_atomic.AddUint64(&mmProduct.afterProductCounter, 1)

	if mmProduct.inspectFuncProduct != nil {
		mmProduct.inspectFuncProduct(ctx, sku)
	}

	mm_params := &ProductServiceClientMockProductParams{ctx, sku}

	// Record call args
	mmProduct.ProductMock.mutex.Lock()
	mmProduct.ProductMock.callArgs = append(mmProduct.ProductMock.callArgs, mm_params)
	mmProduct.ProductMock.mutex.Unlock()

	for _, e := range mmProduct.ProductMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmProduct.ProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmProduct.ProductMock.defaultExpectation.Counter, 1)
		mm_want := mmProduct.ProductMock.defaultExpectation.params
		mm_got := ProductServiceClientMockProductParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmProduct.t.Errorf("ProductServiceClientMock.Product got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmProduct.ProductMock.defaultExpectation.results
		if mm_results == nil {
			mmProduct.t.Fatal("No results are set for the ProductServiceClientMock.Product")
		}
		return (*mm_results).p1, (*mm_results).err
	}
	if mmProduct.funcProduct != nil {
		return mmProduct.funcProduct(ctx, sku)
	}
	mmProduct.t.Fatalf("Unexpected call to ProductServiceClientMock.Product. %v %v", ctx, sku)
	return
}

// ProductAfterCounter returns a count of finished ProductServiceClientMock.Product invocations
func (mmProduct *ProductServiceClientMock) ProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProduct.afterProductCounter)
}

// ProductBeforeCounter returns a count of ProductServiceClientMock.Product invocations
func (mmProduct *ProductServiceClientMock) ProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmProduct.beforeProductCounter)
}

// Calls returns a list of arguments used in each call to ProductServiceClientMock.Product.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmProduct *mProductServiceClientMockProduct) Calls() []*ProductServiceClientMockProductParams {
	mmProduct.mutex.RLock()

	argCopy := make([]*ProductServiceClientMockProductParams, len(mmProduct.callArgs))
	copy(argCopy, mmProduct.callArgs)

	mmProduct.mutex.RUnlock()

	return argCopy
}

// MinimockProductDone returns true if the count of the Product invocations corresponds
// the number of defined expectations
func (m *ProductServiceClientMock) MinimockProductDone() bool {
	for _, e := range m.ProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProductCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProduct != nil && mm_atomic.LoadUint64(&m.afterProductCounter) < 1 {
		return false
	}
	return true
}

// MinimockProductInspect logs each unmet expectation
func (m *ProductServiceClientMock) MinimockProductInspect() {
	for _, e := range m.ProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProductServiceClientMock.Product with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterProductCounter) < 1 {
		if m.ProductMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ProductServiceClientMock.Product")
		} else {
			m.t.Errorf("Expected call to ProductServiceClientMock.Product with params: %#v", *m.ProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcProduct != nil && mm_atomic.LoadUint64(&m.afterProductCounter) < 1 {
		m.t.Error("Expected call to ProductServiceClientMock.Product")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProductServiceClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockProductInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProductServiceClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ProductServiceClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockProductDone()
}