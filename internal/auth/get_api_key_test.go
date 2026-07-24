package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if apiKey != "abc123" {
		t.Errorf("expected api key abc123, got %s", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer abc123")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Error("expected an error for malformed authorization header")
	}
}
