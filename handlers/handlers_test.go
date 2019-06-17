// handlers_test.go
package handlers

import (
	"github.com/golang/mock/gomock"
	"github.com/nammn/node-aggregation/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRedis := mocks.NewMockRedisConnection(mockCtrl)

	handlers := NewHandler(mockRedis)

	handler := http.HandlerFunc(handlers.HealthCheckHandler)
	gomock.InOrder(
		mockRedis.EXPECT().Ping().Return(nil),
	)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUnHealthyCheck(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRedis := mocks.NewMockRedisConnection(mockCtrl)

	handlers := NewHandler(mockRedis)

	handler := http.HandlerFunc(handlers.HealthCheckHandler)
	gomock.InOrder(
		mockRedis.EXPECT().Ping().Return(err),
	)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status == http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
