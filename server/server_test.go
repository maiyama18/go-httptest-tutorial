package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/message", nil)
	rec := httptest.NewRecorder()

	MessageHandler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Error("The status code is not OK")
	}

	fmt.Println("response", rec.Body.String())
}
