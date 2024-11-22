// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package server

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/AskaryanKarine/bmstu-ds-3/pkg/models"
	"github.com/gojuno/minimock/v3"
)

// LoyaltyRepositoryMock implements loyaltyRepository
type LoyaltyRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetByUser          func(ctx context.Context, username string) (l1 models.LoyaltyInfoResponse, err error)
	funcGetByUserOrigin    string
	inspectFuncGetByUser   func(ctx context.Context, username string)
	afterGetByUserCounter  uint64
	beforeGetByUserCounter uint64
	GetByUserMock          mLoyaltyRepositoryMockGetByUser

	funcUpdateByUser          func(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) (err error)
	funcUpdateByUserOrigin    string
	inspectFuncUpdateByUser   func(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse)
	afterUpdateByUserCounter  uint64
	beforeUpdateByUserCounter uint64
	UpdateByUserMock          mLoyaltyRepositoryMockUpdateByUser
}

// NewLoyaltyRepositoryMock returns a mock for loyaltyRepository
func NewLoyaltyRepositoryMock(t minimock.Tester) *LoyaltyRepositoryMock {
	m := &LoyaltyRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetByUserMock = mLoyaltyRepositoryMockGetByUser{mock: m}
	m.GetByUserMock.callArgs = []*LoyaltyRepositoryMockGetByUserParams{}

	m.UpdateByUserMock = mLoyaltyRepositoryMockUpdateByUser{mock: m}
	m.UpdateByUserMock.callArgs = []*LoyaltyRepositoryMockUpdateByUserParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLoyaltyRepositoryMockGetByUser struct {
	optional           bool
	mock               *LoyaltyRepositoryMock
	defaultExpectation *LoyaltyRepositoryMockGetByUserExpectation
	expectations       []*LoyaltyRepositoryMockGetByUserExpectation

	callArgs []*LoyaltyRepositoryMockGetByUserParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// LoyaltyRepositoryMockGetByUserExpectation specifies expectation struct of the loyaltyRepository.GetByUser
type LoyaltyRepositoryMockGetByUserExpectation struct {
	mock               *LoyaltyRepositoryMock
	params             *LoyaltyRepositoryMockGetByUserParams
	paramPtrs          *LoyaltyRepositoryMockGetByUserParamPtrs
	expectationOrigins LoyaltyRepositoryMockGetByUserExpectationOrigins
	results            *LoyaltyRepositoryMockGetByUserResults
	returnOrigin       string
	Counter            uint64
}

// LoyaltyRepositoryMockGetByUserParams contains parameters of the loyaltyRepository.GetByUser
type LoyaltyRepositoryMockGetByUserParams struct {
	ctx      context.Context
	username string
}

// LoyaltyRepositoryMockGetByUserParamPtrs contains pointers to parameters of the loyaltyRepository.GetByUser
type LoyaltyRepositoryMockGetByUserParamPtrs struct {
	ctx      *context.Context
	username *string
}

// LoyaltyRepositoryMockGetByUserResults contains results of the loyaltyRepository.GetByUser
type LoyaltyRepositoryMockGetByUserResults struct {
	l1  models.LoyaltyInfoResponse
	err error
}

// LoyaltyRepositoryMockGetByUserOrigins contains origins of expectations of the loyaltyRepository.GetByUser
type LoyaltyRepositoryMockGetByUserExpectationOrigins struct {
	origin         string
	originCtx      string
	originUsername string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Optional() *mLoyaltyRepositoryMockGetByUser {
	mmGetByUser.optional = true
	return mmGetByUser
}

// Expect sets up expected params for loyaltyRepository.GetByUser
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Expect(ctx context.Context, username string) *mLoyaltyRepositoryMockGetByUser {
	if mmGetByUser.mock.funcGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Set")
	}

	if mmGetByUser.defaultExpectation == nil {
		mmGetByUser.defaultExpectation = &LoyaltyRepositoryMockGetByUserExpectation{}
	}

	if mmGetByUser.defaultExpectation.paramPtrs != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by ExpectParams functions")
	}

	mmGetByUser.defaultExpectation.params = &LoyaltyRepositoryMockGetByUserParams{ctx, username}
	mmGetByUser.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetByUser.expectations {
		if minimock.Equal(e.params, mmGetByUser.defaultExpectation.params) {
			mmGetByUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetByUser.defaultExpectation.params)
		}
	}

	return mmGetByUser
}

// ExpectCtxParam1 sets up expected param ctx for loyaltyRepository.GetByUser
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) ExpectCtxParam1(ctx context.Context) *mLoyaltyRepositoryMockGetByUser {
	if mmGetByUser.mock.funcGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Set")
	}

	if mmGetByUser.defaultExpectation == nil {
		mmGetByUser.defaultExpectation = &LoyaltyRepositoryMockGetByUserExpectation{}
	}

	if mmGetByUser.defaultExpectation.params != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Expect")
	}

	if mmGetByUser.defaultExpectation.paramPtrs == nil {
		mmGetByUser.defaultExpectation.paramPtrs = &LoyaltyRepositoryMockGetByUserParamPtrs{}
	}
	mmGetByUser.defaultExpectation.paramPtrs.ctx = &ctx
	mmGetByUser.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmGetByUser
}

// ExpectUsernameParam2 sets up expected param username for loyaltyRepository.GetByUser
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) ExpectUsernameParam2(username string) *mLoyaltyRepositoryMockGetByUser {
	if mmGetByUser.mock.funcGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Set")
	}

	if mmGetByUser.defaultExpectation == nil {
		mmGetByUser.defaultExpectation = &LoyaltyRepositoryMockGetByUserExpectation{}
	}

	if mmGetByUser.defaultExpectation.params != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Expect")
	}

	if mmGetByUser.defaultExpectation.paramPtrs == nil {
		mmGetByUser.defaultExpectation.paramPtrs = &LoyaltyRepositoryMockGetByUserParamPtrs{}
	}
	mmGetByUser.defaultExpectation.paramPtrs.username = &username
	mmGetByUser.defaultExpectation.expectationOrigins.originUsername = minimock.CallerInfo(1)

	return mmGetByUser
}

// Inspect accepts an inspector function that has same arguments as the loyaltyRepository.GetByUser
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Inspect(f func(ctx context.Context, username string)) *mLoyaltyRepositoryMockGetByUser {
	if mmGetByUser.mock.inspectFuncGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("Inspect function is already set for LoyaltyRepositoryMock.GetByUser")
	}

	mmGetByUser.mock.inspectFuncGetByUser = f

	return mmGetByUser
}

// Return sets up results that will be returned by loyaltyRepository.GetByUser
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Return(l1 models.LoyaltyInfoResponse, err error) *LoyaltyRepositoryMock {
	if mmGetByUser.mock.funcGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Set")
	}

	if mmGetByUser.defaultExpectation == nil {
		mmGetByUser.defaultExpectation = &LoyaltyRepositoryMockGetByUserExpectation{mock: mmGetByUser.mock}
	}
	mmGetByUser.defaultExpectation.results = &LoyaltyRepositoryMockGetByUserResults{l1, err}
	mmGetByUser.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetByUser.mock
}

// Set uses given function f to mock the loyaltyRepository.GetByUser method
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Set(f func(ctx context.Context, username string) (l1 models.LoyaltyInfoResponse, err error)) *LoyaltyRepositoryMock {
	if mmGetByUser.defaultExpectation != nil {
		mmGetByUser.mock.t.Fatalf("Default expectation is already set for the loyaltyRepository.GetByUser method")
	}

	if len(mmGetByUser.expectations) > 0 {
		mmGetByUser.mock.t.Fatalf("Some expectations are already set for the loyaltyRepository.GetByUser method")
	}

	mmGetByUser.mock.funcGetByUser = f
	mmGetByUser.mock.funcGetByUserOrigin = minimock.CallerInfo(1)
	return mmGetByUser.mock
}

// When sets expectation for the loyaltyRepository.GetByUser which will trigger the result defined by the following
// Then helper
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) When(ctx context.Context, username string) *LoyaltyRepositoryMockGetByUserExpectation {
	if mmGetByUser.mock.funcGetByUser != nil {
		mmGetByUser.mock.t.Fatalf("LoyaltyRepositoryMock.GetByUser mock is already set by Set")
	}

	expectation := &LoyaltyRepositoryMockGetByUserExpectation{
		mock:               mmGetByUser.mock,
		params:             &LoyaltyRepositoryMockGetByUserParams{ctx, username},
		expectationOrigins: LoyaltyRepositoryMockGetByUserExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetByUser.expectations = append(mmGetByUser.expectations, expectation)
	return expectation
}

// Then sets up loyaltyRepository.GetByUser return parameters for the expectation previously defined by the When method
func (e *LoyaltyRepositoryMockGetByUserExpectation) Then(l1 models.LoyaltyInfoResponse, err error) *LoyaltyRepositoryMock {
	e.results = &LoyaltyRepositoryMockGetByUserResults{l1, err}
	return e.mock
}

// Times sets number of times loyaltyRepository.GetByUser should be invoked
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Times(n uint64) *mLoyaltyRepositoryMockGetByUser {
	if n == 0 {
		mmGetByUser.mock.t.Fatalf("Times of LoyaltyRepositoryMock.GetByUser mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetByUser.expectedInvocations, n)
	mmGetByUser.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetByUser
}

func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) invocationsDone() bool {
	if len(mmGetByUser.expectations) == 0 && mmGetByUser.defaultExpectation == nil && mmGetByUser.mock.funcGetByUser == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetByUser.mock.afterGetByUserCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetByUser.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetByUser implements loyaltyRepository
func (mmGetByUser *LoyaltyRepositoryMock) GetByUser(ctx context.Context, username string) (l1 models.LoyaltyInfoResponse, err error) {
	mm_atomic.AddUint64(&mmGetByUser.beforeGetByUserCounter, 1)
	defer mm_atomic.AddUint64(&mmGetByUser.afterGetByUserCounter, 1)

	mmGetByUser.t.Helper()

	if mmGetByUser.inspectFuncGetByUser != nil {
		mmGetByUser.inspectFuncGetByUser(ctx, username)
	}

	mm_params := LoyaltyRepositoryMockGetByUserParams{ctx, username}

	// Record call args
	mmGetByUser.GetByUserMock.mutex.Lock()
	mmGetByUser.GetByUserMock.callArgs = append(mmGetByUser.GetByUserMock.callArgs, &mm_params)
	mmGetByUser.GetByUserMock.mutex.Unlock()

	for _, e := range mmGetByUser.GetByUserMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.l1, e.results.err
		}
	}

	if mmGetByUser.GetByUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetByUser.GetByUserMock.defaultExpectation.Counter, 1)
		mm_want := mmGetByUser.GetByUserMock.defaultExpectation.params
		mm_want_ptrs := mmGetByUser.GetByUserMock.defaultExpectation.paramPtrs

		mm_got := LoyaltyRepositoryMockGetByUserParams{ctx, username}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGetByUser.t.Errorf("LoyaltyRepositoryMock.GetByUser got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetByUser.GetByUserMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.username != nil && !minimock.Equal(*mm_want_ptrs.username, mm_got.username) {
				mmGetByUser.t.Errorf("LoyaltyRepositoryMock.GetByUser got unexpected parameter username, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetByUser.GetByUserMock.defaultExpectation.expectationOrigins.originUsername, *mm_want_ptrs.username, mm_got.username, minimock.Diff(*mm_want_ptrs.username, mm_got.username))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetByUser.t.Errorf("LoyaltyRepositoryMock.GetByUser got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetByUser.GetByUserMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetByUser.GetByUserMock.defaultExpectation.results
		if mm_results == nil {
			mmGetByUser.t.Fatal("No results are set for the LoyaltyRepositoryMock.GetByUser")
		}
		return (*mm_results).l1, (*mm_results).err
	}
	if mmGetByUser.funcGetByUser != nil {
		return mmGetByUser.funcGetByUser(ctx, username)
	}
	mmGetByUser.t.Fatalf("Unexpected call to LoyaltyRepositoryMock.GetByUser. %v %v", ctx, username)
	return
}

// GetByUserAfterCounter returns a count of finished LoyaltyRepositoryMock.GetByUser invocations
func (mmGetByUser *LoyaltyRepositoryMock) GetByUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByUser.afterGetByUserCounter)
}

// GetByUserBeforeCounter returns a count of LoyaltyRepositoryMock.GetByUser invocations
func (mmGetByUser *LoyaltyRepositoryMock) GetByUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByUser.beforeGetByUserCounter)
}

// Calls returns a list of arguments used in each call to LoyaltyRepositoryMock.GetByUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetByUser *mLoyaltyRepositoryMockGetByUser) Calls() []*LoyaltyRepositoryMockGetByUserParams {
	mmGetByUser.mutex.RLock()

	argCopy := make([]*LoyaltyRepositoryMockGetByUserParams, len(mmGetByUser.callArgs))
	copy(argCopy, mmGetByUser.callArgs)

	mmGetByUser.mutex.RUnlock()

	return argCopy
}

// MinimockGetByUserDone returns true if the count of the GetByUser invocations corresponds
// the number of defined expectations
func (m *LoyaltyRepositoryMock) MinimockGetByUserDone() bool {
	if m.GetByUserMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetByUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetByUserMock.invocationsDone()
}

// MinimockGetByUserInspect logs each unmet expectation
func (m *LoyaltyRepositoryMock) MinimockGetByUserInspect() {
	for _, e := range m.GetByUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.GetByUser at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetByUserCounter := mm_atomic.LoadUint64(&m.afterGetByUserCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetByUserMock.defaultExpectation != nil && afterGetByUserCounter < 1 {
		if m.GetByUserMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.GetByUser at\n%s", m.GetByUserMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.GetByUser at\n%s with params: %#v", m.GetByUserMock.defaultExpectation.expectationOrigins.origin, *m.GetByUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByUser != nil && afterGetByUserCounter < 1 {
		m.t.Errorf("Expected call to LoyaltyRepositoryMock.GetByUser at\n%s", m.funcGetByUserOrigin)
	}

	if !m.GetByUserMock.invocationsDone() && afterGetByUserCounter > 0 {
		m.t.Errorf("Expected %d calls to LoyaltyRepositoryMock.GetByUser at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetByUserMock.expectedInvocations), m.GetByUserMock.expectedInvocationsOrigin, afterGetByUserCounter)
	}
}

type mLoyaltyRepositoryMockUpdateByUser struct {
	optional           bool
	mock               *LoyaltyRepositoryMock
	defaultExpectation *LoyaltyRepositoryMockUpdateByUserExpectation
	expectations       []*LoyaltyRepositoryMockUpdateByUserExpectation

	callArgs []*LoyaltyRepositoryMockUpdateByUserParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// LoyaltyRepositoryMockUpdateByUserExpectation specifies expectation struct of the loyaltyRepository.UpdateByUser
type LoyaltyRepositoryMockUpdateByUserExpectation struct {
	mock               *LoyaltyRepositoryMock
	params             *LoyaltyRepositoryMockUpdateByUserParams
	paramPtrs          *LoyaltyRepositoryMockUpdateByUserParamPtrs
	expectationOrigins LoyaltyRepositoryMockUpdateByUserExpectationOrigins
	results            *LoyaltyRepositoryMockUpdateByUserResults
	returnOrigin       string
	Counter            uint64
}

// LoyaltyRepositoryMockUpdateByUserParams contains parameters of the loyaltyRepository.UpdateByUser
type LoyaltyRepositoryMockUpdateByUserParams struct {
	ctx          context.Context
	username     string
	usersLoyalty models.LoyaltyInfoResponse
}

// LoyaltyRepositoryMockUpdateByUserParamPtrs contains pointers to parameters of the loyaltyRepository.UpdateByUser
type LoyaltyRepositoryMockUpdateByUserParamPtrs struct {
	ctx          *context.Context
	username     *string
	usersLoyalty *models.LoyaltyInfoResponse
}

// LoyaltyRepositoryMockUpdateByUserResults contains results of the loyaltyRepository.UpdateByUser
type LoyaltyRepositoryMockUpdateByUserResults struct {
	err error
}

// LoyaltyRepositoryMockUpdateByUserOrigins contains origins of expectations of the loyaltyRepository.UpdateByUser
type LoyaltyRepositoryMockUpdateByUserExpectationOrigins struct {
	origin             string
	originCtx          string
	originUsername     string
	originUsersLoyalty string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Optional() *mLoyaltyRepositoryMockUpdateByUser {
	mmUpdateByUser.optional = true
	return mmUpdateByUser
}

// Expect sets up expected params for loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Expect(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) *mLoyaltyRepositoryMockUpdateByUser {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	if mmUpdateByUser.defaultExpectation == nil {
		mmUpdateByUser.defaultExpectation = &LoyaltyRepositoryMockUpdateByUserExpectation{}
	}

	if mmUpdateByUser.defaultExpectation.paramPtrs != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by ExpectParams functions")
	}

	mmUpdateByUser.defaultExpectation.params = &LoyaltyRepositoryMockUpdateByUserParams{ctx, username, usersLoyalty}
	mmUpdateByUser.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmUpdateByUser.expectations {
		if minimock.Equal(e.params, mmUpdateByUser.defaultExpectation.params) {
			mmUpdateByUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateByUser.defaultExpectation.params)
		}
	}

	return mmUpdateByUser
}

// ExpectCtxParam1 sets up expected param ctx for loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) ExpectCtxParam1(ctx context.Context) *mLoyaltyRepositoryMockUpdateByUser {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	if mmUpdateByUser.defaultExpectation == nil {
		mmUpdateByUser.defaultExpectation = &LoyaltyRepositoryMockUpdateByUserExpectation{}
	}

	if mmUpdateByUser.defaultExpectation.params != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Expect")
	}

	if mmUpdateByUser.defaultExpectation.paramPtrs == nil {
		mmUpdateByUser.defaultExpectation.paramPtrs = &LoyaltyRepositoryMockUpdateByUserParamPtrs{}
	}
	mmUpdateByUser.defaultExpectation.paramPtrs.ctx = &ctx
	mmUpdateByUser.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmUpdateByUser
}

// ExpectUsernameParam2 sets up expected param username for loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) ExpectUsernameParam2(username string) *mLoyaltyRepositoryMockUpdateByUser {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	if mmUpdateByUser.defaultExpectation == nil {
		mmUpdateByUser.defaultExpectation = &LoyaltyRepositoryMockUpdateByUserExpectation{}
	}

	if mmUpdateByUser.defaultExpectation.params != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Expect")
	}

	if mmUpdateByUser.defaultExpectation.paramPtrs == nil {
		mmUpdateByUser.defaultExpectation.paramPtrs = &LoyaltyRepositoryMockUpdateByUserParamPtrs{}
	}
	mmUpdateByUser.defaultExpectation.paramPtrs.username = &username
	mmUpdateByUser.defaultExpectation.expectationOrigins.originUsername = minimock.CallerInfo(1)

	return mmUpdateByUser
}

// ExpectUsersLoyaltyParam3 sets up expected param usersLoyalty for loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) ExpectUsersLoyaltyParam3(usersLoyalty models.LoyaltyInfoResponse) *mLoyaltyRepositoryMockUpdateByUser {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	if mmUpdateByUser.defaultExpectation == nil {
		mmUpdateByUser.defaultExpectation = &LoyaltyRepositoryMockUpdateByUserExpectation{}
	}

	if mmUpdateByUser.defaultExpectation.params != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Expect")
	}

	if mmUpdateByUser.defaultExpectation.paramPtrs == nil {
		mmUpdateByUser.defaultExpectation.paramPtrs = &LoyaltyRepositoryMockUpdateByUserParamPtrs{}
	}
	mmUpdateByUser.defaultExpectation.paramPtrs.usersLoyalty = &usersLoyalty
	mmUpdateByUser.defaultExpectation.expectationOrigins.originUsersLoyalty = minimock.CallerInfo(1)

	return mmUpdateByUser
}

// Inspect accepts an inspector function that has same arguments as the loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Inspect(f func(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse)) *mLoyaltyRepositoryMockUpdateByUser {
	if mmUpdateByUser.mock.inspectFuncUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("Inspect function is already set for LoyaltyRepositoryMock.UpdateByUser")
	}

	mmUpdateByUser.mock.inspectFuncUpdateByUser = f

	return mmUpdateByUser
}

// Return sets up results that will be returned by loyaltyRepository.UpdateByUser
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Return(err error) *LoyaltyRepositoryMock {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	if mmUpdateByUser.defaultExpectation == nil {
		mmUpdateByUser.defaultExpectation = &LoyaltyRepositoryMockUpdateByUserExpectation{mock: mmUpdateByUser.mock}
	}
	mmUpdateByUser.defaultExpectation.results = &LoyaltyRepositoryMockUpdateByUserResults{err}
	mmUpdateByUser.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmUpdateByUser.mock
}

// Set uses given function f to mock the loyaltyRepository.UpdateByUser method
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Set(f func(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) (err error)) *LoyaltyRepositoryMock {
	if mmUpdateByUser.defaultExpectation != nil {
		mmUpdateByUser.mock.t.Fatalf("Default expectation is already set for the loyaltyRepository.UpdateByUser method")
	}

	if len(mmUpdateByUser.expectations) > 0 {
		mmUpdateByUser.mock.t.Fatalf("Some expectations are already set for the loyaltyRepository.UpdateByUser method")
	}

	mmUpdateByUser.mock.funcUpdateByUser = f
	mmUpdateByUser.mock.funcUpdateByUserOrigin = minimock.CallerInfo(1)
	return mmUpdateByUser.mock
}

// When sets expectation for the loyaltyRepository.UpdateByUser which will trigger the result defined by the following
// Then helper
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) When(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) *LoyaltyRepositoryMockUpdateByUserExpectation {
	if mmUpdateByUser.mock.funcUpdateByUser != nil {
		mmUpdateByUser.mock.t.Fatalf("LoyaltyRepositoryMock.UpdateByUser mock is already set by Set")
	}

	expectation := &LoyaltyRepositoryMockUpdateByUserExpectation{
		mock:               mmUpdateByUser.mock,
		params:             &LoyaltyRepositoryMockUpdateByUserParams{ctx, username, usersLoyalty},
		expectationOrigins: LoyaltyRepositoryMockUpdateByUserExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmUpdateByUser.expectations = append(mmUpdateByUser.expectations, expectation)
	return expectation
}

// Then sets up loyaltyRepository.UpdateByUser return parameters for the expectation previously defined by the When method
func (e *LoyaltyRepositoryMockUpdateByUserExpectation) Then(err error) *LoyaltyRepositoryMock {
	e.results = &LoyaltyRepositoryMockUpdateByUserResults{err}
	return e.mock
}

// Times sets number of times loyaltyRepository.UpdateByUser should be invoked
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Times(n uint64) *mLoyaltyRepositoryMockUpdateByUser {
	if n == 0 {
		mmUpdateByUser.mock.t.Fatalf("Times of LoyaltyRepositoryMock.UpdateByUser mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmUpdateByUser.expectedInvocations, n)
	mmUpdateByUser.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmUpdateByUser
}

func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) invocationsDone() bool {
	if len(mmUpdateByUser.expectations) == 0 && mmUpdateByUser.defaultExpectation == nil && mmUpdateByUser.mock.funcUpdateByUser == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmUpdateByUser.mock.afterUpdateByUserCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmUpdateByUser.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// UpdateByUser implements loyaltyRepository
func (mmUpdateByUser *LoyaltyRepositoryMock) UpdateByUser(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) (err error) {
	mm_atomic.AddUint64(&mmUpdateByUser.beforeUpdateByUserCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateByUser.afterUpdateByUserCounter, 1)

	mmUpdateByUser.t.Helper()

	if mmUpdateByUser.inspectFuncUpdateByUser != nil {
		mmUpdateByUser.inspectFuncUpdateByUser(ctx, username, usersLoyalty)
	}

	mm_params := LoyaltyRepositoryMockUpdateByUserParams{ctx, username, usersLoyalty}

	// Record call args
	mmUpdateByUser.UpdateByUserMock.mutex.Lock()
	mmUpdateByUser.UpdateByUserMock.callArgs = append(mmUpdateByUser.UpdateByUserMock.callArgs, &mm_params)
	mmUpdateByUser.UpdateByUserMock.mutex.Unlock()

	for _, e := range mmUpdateByUser.UpdateByUserMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateByUser.UpdateByUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateByUser.UpdateByUserMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateByUser.UpdateByUserMock.defaultExpectation.params
		mm_want_ptrs := mmUpdateByUser.UpdateByUserMock.defaultExpectation.paramPtrs

		mm_got := LoyaltyRepositoryMockUpdateByUserParams{ctx, username, usersLoyalty}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmUpdateByUser.t.Errorf("LoyaltyRepositoryMock.UpdateByUser got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmUpdateByUser.UpdateByUserMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.username != nil && !minimock.Equal(*mm_want_ptrs.username, mm_got.username) {
				mmUpdateByUser.t.Errorf("LoyaltyRepositoryMock.UpdateByUser got unexpected parameter username, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmUpdateByUser.UpdateByUserMock.defaultExpectation.expectationOrigins.originUsername, *mm_want_ptrs.username, mm_got.username, minimock.Diff(*mm_want_ptrs.username, mm_got.username))
			}

			if mm_want_ptrs.usersLoyalty != nil && !minimock.Equal(*mm_want_ptrs.usersLoyalty, mm_got.usersLoyalty) {
				mmUpdateByUser.t.Errorf("LoyaltyRepositoryMock.UpdateByUser got unexpected parameter usersLoyalty, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmUpdateByUser.UpdateByUserMock.defaultExpectation.expectationOrigins.originUsersLoyalty, *mm_want_ptrs.usersLoyalty, mm_got.usersLoyalty, minimock.Diff(*mm_want_ptrs.usersLoyalty, mm_got.usersLoyalty))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateByUser.t.Errorf("LoyaltyRepositoryMock.UpdateByUser got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmUpdateByUser.UpdateByUserMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateByUser.UpdateByUserMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateByUser.t.Fatal("No results are set for the LoyaltyRepositoryMock.UpdateByUser")
		}
		return (*mm_results).err
	}
	if mmUpdateByUser.funcUpdateByUser != nil {
		return mmUpdateByUser.funcUpdateByUser(ctx, username, usersLoyalty)
	}
	mmUpdateByUser.t.Fatalf("Unexpected call to LoyaltyRepositoryMock.UpdateByUser. %v %v %v", ctx, username, usersLoyalty)
	return
}

// UpdateByUserAfterCounter returns a count of finished LoyaltyRepositoryMock.UpdateByUser invocations
func (mmUpdateByUser *LoyaltyRepositoryMock) UpdateByUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateByUser.afterUpdateByUserCounter)
}

// UpdateByUserBeforeCounter returns a count of LoyaltyRepositoryMock.UpdateByUser invocations
func (mmUpdateByUser *LoyaltyRepositoryMock) UpdateByUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateByUser.beforeUpdateByUserCounter)
}

// Calls returns a list of arguments used in each call to LoyaltyRepositoryMock.UpdateByUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateByUser *mLoyaltyRepositoryMockUpdateByUser) Calls() []*LoyaltyRepositoryMockUpdateByUserParams {
	mmUpdateByUser.mutex.RLock()

	argCopy := make([]*LoyaltyRepositoryMockUpdateByUserParams, len(mmUpdateByUser.callArgs))
	copy(argCopy, mmUpdateByUser.callArgs)

	mmUpdateByUser.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateByUserDone returns true if the count of the UpdateByUser invocations corresponds
// the number of defined expectations
func (m *LoyaltyRepositoryMock) MinimockUpdateByUserDone() bool {
	if m.UpdateByUserMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.UpdateByUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.UpdateByUserMock.invocationsDone()
}

// MinimockUpdateByUserInspect logs each unmet expectation
func (m *LoyaltyRepositoryMock) MinimockUpdateByUserInspect() {
	for _, e := range m.UpdateByUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.UpdateByUser at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterUpdateByUserCounter := mm_atomic.LoadUint64(&m.afterUpdateByUserCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateByUserMock.defaultExpectation != nil && afterUpdateByUserCounter < 1 {
		if m.UpdateByUserMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.UpdateByUser at\n%s", m.UpdateByUserMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to LoyaltyRepositoryMock.UpdateByUser at\n%s with params: %#v", m.UpdateByUserMock.defaultExpectation.expectationOrigins.origin, *m.UpdateByUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateByUser != nil && afterUpdateByUserCounter < 1 {
		m.t.Errorf("Expected call to LoyaltyRepositoryMock.UpdateByUser at\n%s", m.funcUpdateByUserOrigin)
	}

	if !m.UpdateByUserMock.invocationsDone() && afterUpdateByUserCounter > 0 {
		m.t.Errorf("Expected %d calls to LoyaltyRepositoryMock.UpdateByUser at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.UpdateByUserMock.expectedInvocations), m.UpdateByUserMock.expectedInvocationsOrigin, afterUpdateByUserCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LoyaltyRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetByUserInspect()

			m.MinimockUpdateByUserInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LoyaltyRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *LoyaltyRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetByUserDone() &&
		m.MinimockUpdateByUserDone()
}
