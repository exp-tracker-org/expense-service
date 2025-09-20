package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRouterSetup(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/expenses", getAllExpenses).Methods("GET")

	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode == 0 {
		t.Errorf("router did not return a valid response")
	}
}
