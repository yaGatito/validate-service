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
	// Mock request body
	requestBody := models.ValidateRequest{
		CardNumber:          "1234567890123456",
		CardYearExpiration:  2024,
		CardMonthExpiration: 3,
	}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/validate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Mock response writer
	w := &MockResponseWriter{}

	// Call the handler function
	handlers.ValidateHandler(w, req)

	// Check response status code
	if w.statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.statusCode)
	}

	// Check response body (decode JSON)
	var response models.ValidateResponse
	err := json.Unmarshal(w.body, &response)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Add more assertions based on your test cases
	// For example:
	if !response.Valid {
		t.Errorf("Expected valid=true in response, got valid=false")
	}
}
