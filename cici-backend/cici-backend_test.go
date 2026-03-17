package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler_ResponseBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	resp := w.Result()
	body := w.Body.String()

	const expected = "hello, my name is cici~~i'm 4 years old\n"
	if body != expected {
		t.Errorf("unexpected body: got %q, want %q", body, expected)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestHelloHandler_StatusCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", w.Code)
	}
}

// Regression test: old response string must no longer be returned.
func TestHelloHandler_OldResponseNotPresent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	oldResponse := "hello, my name is cici~~\n"
	if w.Body.String() == oldResponse {
		t.Errorf("handler returned old response %q; expected updated greeting", oldResponse)
	}
}

// Boundary: response must include the age information added in this PR.
func TestHelloHandler_ContainsAgeInfo(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "4 years old") {
		t.Errorf("expected body to contain age info %q, got %q", "4 years old", body)
	}
}

// Negative: response must not contain extra unexpected content beyond the greeting.
func TestHelloHandler_ExactBodyMatch(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	body := w.Body.String()
	const expected = "hello, my name is cici~~i'm 4 years old\n"
	if len(body) != len(expected) {
		t.Errorf("body length mismatch: got %d bytes (%q), want %d bytes (%q)", len(body), body, len(expected), expected)
	}
}