// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"sync"
	"testing"

	"xenapi"
)

type fakeXAPI struct {
	server *httptest.Server

	results map[string]interface{}

	mu sync.Mutex
	// calls holds each call's arguments by method, without the leading session reference.
	calls map[string][][]interface{}
}

func newFakeXAPI(t *testing.T, results map[string]interface{}) *fakeXAPI {
	t.Helper()

	f := &fakeXAPI{
		results: results,
		calls:   make(map[string][][]interface{}),
	}
	f.server = httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(f.server.Close)

	return f
}

func (f *fakeXAPI) serve(w http.ResponseWriter, r *http.Request) {
	var req xenapi.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params, _ := req.Params.([]interface{})
	if len(params) > 0 {
		params = params[1:]
	}

	f.mu.Lock()
	f.calls[req.Method] = append(f.calls[req.Method], params)
	result, known := f.results[req.Method]
	f.mu.Unlock()

	resp := xenapi.Response{JSONRPC: "2.0", ID: req.ID}
	if !known {
		resp.Error = &xenapi.ResponseError{Code: 1, Message: "MESSAGE_METHOD_UNKNOWN: " + req.Method}
	} else {
		switch result := result.(type) {
		case error:
			resp.Error = &xenapi.ResponseError{Code: 1, Message: result.Error()}
		case func([]interface{}) interface{}:
			switch answer := result(params).(type) {
			case error:
				resp.Error = &xenapi.ResponseError{Code: 1, Message: answer.Error()}
			default:
				resp.Result = answer
			}
		default:
			resp.Result = result
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func (f *fakeXAPI) conn() *Connection {
	return &Connection{session: xenapi.NewSession(&xenapi.ClientOpts{URL: f.server.URL})}
}

func (f *fakeXAPI) methodsCalled() []string {
	f.mu.Lock()
	defer f.mu.Unlock()

	methods := make([]string, 0, len(f.calls))
	for method := range f.calls {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	return methods
}

func (f *fakeXAPI) wasCalled(method string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	return len(f.calls[method]) > 0
}

func (f *fakeXAPI) allArgsOf(method string) [][]interface{} {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.calls[method]
}

func (f *fakeXAPI) argsOf(t *testing.T, method string) []interface{} {
	t.Helper()

	f.mu.Lock()
	defer f.mu.Unlock()

	calls := f.calls[method]
	if len(calls) != 1 {
		t.Fatalf("expected exactly one call to %s, got %d", method, len(calls))
	}

	return calls[0]
}
