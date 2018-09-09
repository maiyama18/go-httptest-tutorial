package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ymr-39/go-httptest-tutorial"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		msg := ghtt.Message{Body: "hello", From: "Tokyo"}

		msgJson, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		w.Write(msgJson)
	})

	log.Fatal(http.ListenAndServe(":5000", mux))
}
