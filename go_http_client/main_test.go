package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := `{"id": 1, "title": "Test Todo", "completed": false}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatusCode)
		fmt.Fprintln(w, expectedBody)
	}))
	defer server.Close()

	statusCode, body, err := makeRequest(server.URL)

	if err != nil {
		t.Fatalf("makeRequest returned an error: %v", err)
	}

	if statusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, statusCode)
	}

	// Trim newline character from body if present, as Fprintln adds one
	trimmedBody := body
	if len(body) > 0 && body[len(body)-1] == '\n' {
		trimmedBody = body[:len(body)-1]
	}


	if trimmedBody != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, trimmedBody)
	}
}
