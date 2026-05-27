package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandlerReturnsHelloWorld(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	routes().ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	expected := "Hello, World!"
	if string(body) != expected {
		t.Fatalf("expected body %q, got %q", expected, string(body))
	}
}

func TestUnknownPathReturnsNotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	routes().ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status code %d, got %d", http.StatusNotFound, response.StatusCode)
	}
}

func TestGetPortReturnsDefaultPort(t *testing.T) {
	t.Setenv("PORT", "")

	port := getPort()

	if port != defaultPort {
		t.Fatalf("expected default port %q, got %q", defaultPort, port)
	}
}
