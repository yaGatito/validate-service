package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"validate-service/src/handlers"
	"validate-service/src/models"
)

// MockResponseWriter implements http.ResponseWriter for testing purposes
type MockResponseWriter struct {
	statusCode int
	body       []byte
}

func (m *MockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

func (m *MockResponseWriter) Write(body []byte) (int, error) {
	m.body = body
	return len(body), nil
}

func TestValidateHandlerValidCard(t *testing.T) {
	// Setup
	requestBody := models.ValidateRequest{
		CardNumber:          "1234567890123456",
		CardYearExpiration:  2024,
		CardMonthExpiration: 3,
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("GET", "/validate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := &MockResponseWriter{}

	// Execution
	handlers.HandleValidatePath(w, req)

	// Assertion
	if w.statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.statusCode)
	}

	// Check response body (decode JSON)
	var response models.ValidateResponse
	err := json.Unmarshal(w.body, &response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if !response.Valid {
		t.Errorf("Expected valid=true in response, got valid=false")
	}
}

func TestValidateHandlerInvalidCardMonth(t *testing.T) {
	// Setup
	requestBody := models.ValidateRequest{
		CardNumber:          "1234567890123456",
		CardYearExpiration:  2024,
		CardMonthExpiration: 23,
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("GET", "/validate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := &MockResponseWriter{}

	// Execution
	handlers.HandleValidatePath(w, req)

	// Assertion
	if w.statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.statusCode)
	}

	var response models.ValidateResponse
	err := json.Unmarshal(w.body, &response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if response.Valid {
		t.Errorf("Expected valid=false in response, got valid=true")
	}
}

func TestValidateHandlerWrongHTTPMethod(t *testing.T) {
	// Setup
	requestBody := models.ValidateRequest{
		CardNumber:          "1234567890123456",
		CardYearExpiration:  2024,
		CardMonthExpiration: 3,
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/validate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := &MockResponseWriter{}

	// Execution
	handlers.HandleValidatePath(w, req)

	// Assertion
	if w.statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.statusCode)
	}

	if len(w.body) != 1 {
		t.Errorf("Body length expected to be 0. But was: %d", len(w.body))
	}
}
