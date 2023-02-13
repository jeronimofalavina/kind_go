package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApp(t *testing.T) {
	req, err := http.NewRequest("GET", "/app", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: resulted %v expected %v",
			status, http.StatusOK)
	}
}
