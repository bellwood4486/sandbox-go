package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func newArgs(method, target string, body io.Reader) args {
	return args{
		httptest.NewRecorder(),
		httptest.NewRequest(method, target, body),
	}
}

func Test_getOneEvent(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			"ok",
			newArgs(http.MethodGet, "/events/1", nil),
			http.StatusOK,
			"./testdata/main/Test_getOneEvent/ok.golden",
		},
		{
			"not found",
			newArgs(http.MethodGet, "/events/100", nil),
			http.StatusNotFound,
			"./testdata/main/Test_getOneEvent/not_found.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRouter().ServeHTTP(tt.args.w, tt.args.r)
			AssertResponse(t, tt.args.w.Result(), tt.wantCode, tt.wantBody)
		})
	}
}

func Test_createEvent(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			"created",
			newArgs(http.MethodPost, "/event", GetReaderFromTestFile(t, "./testdata/main/Test_createEvent/created.json")),
			http.StatusCreated,
			"./testdata/main/Test_createEvent/created.golden",
		},
		{
			"invalid json",
			newArgs(http.MethodPost, "/event", GetReaderFromTestFile(t, "./testdata/main/Test_createEvent/invalid_json.json")),
			http.StatusBadRequest,
			"./testdata/main/Test_createEvent/invalid_json.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRouter().ServeHTTP(tt.args.w, tt.args.r)
			AssertResponse(t, tt.args.w.Result(), tt.wantCode, tt.wantBody)
		})
	}
}

func Test_updateEvent(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			"ok",
			newArgs(http.MethodPatch, "/events/1", GetReaderFromTestFile(t, "./testdata/main/Test_updateEvent/ok.json")),
			http.StatusOK,
			"./testdata/main/Test_updateEvent/ok.golden",
		},
		{
			"invalid json",
			newArgs(http.MethodPatch, "/events/1", GetReaderFromTestFile(t, "./testdata/main/Test_updateEvent/invalid_json.json")),
			http.StatusBadRequest,
			"./testdata/main/Test_updateEvent/invalid_json.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRouter().ServeHTTP(tt.args.w, tt.args.r)
			AssertResponse(t, tt.args.w.Result(), tt.wantCode, tt.wantBody)
		})
	}
}

func Test_deleteEvent(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			"no content",
			newArgs(http.MethodDelete, "/events/1", nil),
			http.StatusNoContent,
			"",
		},
		{
			"not found",
			newArgs(http.MethodDelete, "/events/100", nil),
			http.StatusNotFound,
			"./testdata/main/Test_deleteEvent/not_found.golden",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newRouter().ServeHTTP(tt.args.w, tt.args.r)
			AssertResponse(t, tt.args.w.Result(), tt.wantCode, tt.wantBody)
		})
	}
}

// AssertResponse assert response header and body.
func AssertResponse(t *testing.T, res *http.Response, code int, path string) {
	t.Helper()

	AssertResponseHeader(t, res, code)
	if path == "" {
		AssertResponseBodyEmpty(t, res)
	} else {
		AssertResponseBodyWithFile(t, res, path)
	}
}

// AssertResponseHeader assert response header.
func AssertResponseHeader(t *testing.T, res *http.Response, code int) {
	t.Helper()

	// ???????????????????????????????????????
	if code != res.StatusCode {
		t.Errorf("expected status code is '%d',\n but actual given code is '%d'", code, res.StatusCode)
	}
	// Content-Type???????????????
	if expected := "application/json; charset=utf-8"; res.Header.Get("Content-Type") != expected {
		t.Errorf("unexpected response Content-Type,\n expected: %#v,\n but given #%v", expected, res.Header.Get("Content-Type"))
	}
}

// AssertResponseBodyWithFile assert response body with test file.
func AssertResponseBodyWithFile(t *testing.T, res *http.Response, path string) {
	t.Helper()

	rs := GetStringFromTestFile(t, path)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error by ioutil.ReadAll() '%#v'", err)
	}
	var actual bytes.Buffer
	err = json.Indent(&actual, body, "", "  ")
	if err != nil {
		t.Fatalf("unexpected error by json.Indent '%#v'", err)
	}
	assert.JSONEq(t, rs, actual.String())
}

func AssertResponseBodyEmpty(t *testing.T, res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error by ioutil.ReadAll() '%#v'", err)
	}
	assert.Empty(t, string(body))
}

// GetStringFromTestFile get string from test file.
func GetStringFromTestFile(t *testing.T, path string) string {
	t.Helper()

	bt, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}
	return string(bt)
}

// GetReaderFromTestFile returns a reader of test file.
func GetReaderFromTestFile(t *testing.T, path string) io.Reader {
	return strings.NewReader(GetStringFromTestFile(t, path))
}
