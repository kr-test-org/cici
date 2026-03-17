package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloResponseBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	expected := "hello, my name is cici~~i'm 4 years old\n"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}

func TestHelloContainsAgeInfo(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	body := w.Body.String()
	if !strings.Contains(body, "4 years old") {
		t.Errorf("expected body to contain age info %q, got %q", "4 years old", body)
	}
}

// Regression test: ensure old response string is no longer returned.
func TestHelloDoesNotReturnOldResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	hello(w, req)

	body := w.Body.String()
	oldResponse := "hello, my name is cici~~\n"
	if body == oldResponse {
		t.Errorf("hello() still returns old response %q; expected updated response", oldResponse)
	}
}

// Boundary: hello() must be idempotent across multiple calls.
func TestHelloIdempotent(t *testing.T) {
	expected := "hello, my name is cici~~i'm 4 years old\n"

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		hello(w, req)

		body := w.Body.String()
		if body != expected {
			t.Errorf("call %d: expected body %q, got %q", i+1, expected, body)
		}
	}
}