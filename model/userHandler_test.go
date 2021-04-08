package model_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nstoker/gorocktrack/model"
)

func TestUserIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(model.UserIndex)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, expected %v",
			rr.Code, http.StatusOK)
	}
	// How to check content?
}
