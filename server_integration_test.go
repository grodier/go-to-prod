package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckEndpoint(t *testing.T) {

	t.Run("healthcheck handler", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/healthcheck/", nil)
		response := httptest.NewRecorder()

		healthcheckHandler(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %d, want %d\n", got, want)
		}
	})

	t.Run("healtcheck integration", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(healthcheckHandler))
		defer server.Close()

		resp, err := http.Get(server.URL + "/healthcheck/")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
		}
	})
}
