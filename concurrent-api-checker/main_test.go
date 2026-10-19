package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckURLSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := server.Client()
	result := checkURL(client, server.URL)

	if result.Err != nil {
		t.Fatalf("expected no error, got %v", result.Err)
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, result.StatusCode)
	}

	if result.URL != server.URL {
		t.Errorf("expected URL %s, got %s", server.URL, result.URL)
	}
}

func TestCheckURLNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := server.Client()
	result := checkURL(client, server.URL)

	if result.Err != nil {
		t.Fatalf("expected no request error, got %v", result.Err)
	}

	if result.StatusCode != http.StatusNotFound {
		t.Errorf("expected status %d, got %d",
			http.StatusNotFound, result.StatusCode)
	}
}

func TestCheckURLRequestError(t *testing.T) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	result := checkURL(client, "://invalid-url")

	if result.Err == nil {
		t.Fatal("expected a request error, got nil")
	}

	if result.URL != "://invalid-url" {
		t.Errorf("expected URL %q, got %q", "://invalid-url", result.URL)
	}
}
