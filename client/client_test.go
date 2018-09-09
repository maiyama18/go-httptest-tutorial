package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ymr-39/go-httptest-tutorial"
)

var fixMsg ghtt.Message
var fixMsgJson []byte

func init() {
	f, _ := os.Open("./testdata/msg.json")
	fixMsgJson, _ = ioutil.ReadAll(f)

	json.Unmarshal(fixMsgJson, &fixMsg)
}

func TestGetMessage(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fixMsgJson)
	})
	mockServer := httptest.NewServer(mockHandler)
	defer mockServer.Close()

	msg, err := GetMessage(mockServer.URL)
	if err != nil {
		t.Error(err)
	}

	if msg != fixMsg {
		t.Error("The message got is not equal to fixture")
	}
}

func TestGetMessageFail(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	mockServer := httptest.NewServer(mockHandler)
	defer mockServer.Close()

	_, err := GetMessage(mockServer.URL)
	if err == nil {
		t.Error(err)
	}
}
