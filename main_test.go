package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateRoute(t *testing.T) {
	// Given a started server
	r := SetupRouter()

	// When a request is made without any from and to params
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/rate", nil)
	r.ServeHTTP(w, req)
	// There is a 400 error
	assert.Equal(t, 400, w.Code)

	// When a request is made with only a `from` param
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/rate?from=BTC", nil)
	r.ServeHTTP(w, req)
	// There is a 400 error
	assert.Equal(t, 400, w.Code)

	// When a request is made with only a `to` param
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/rate?to=BTC", nil)
	r.ServeHTTP(w, req)
	// There is a 400 error
	assert.Equal(t, 400, w.Code)

	// When a request is made with both `from` and `to` params
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/rate?from=BTC&to=USD", nil)
	r.ServeHTTP(w, req)
	// There is a 200 success
	assert.Equal(t, 200, w.Code)

	// When a request is made with invalid `from` and `to` params
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/rate?from=INVALID&to=INVALID", nil)
	r.ServeHTTP(w, req)
	// There is a 500 error
	assert.Equal(t, 500, w.Code)

}
