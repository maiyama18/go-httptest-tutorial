package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ghtt "github.com/ymr-39/go-httptest-tutorial"
)

func main() {
	msg, err := GetMessage("http://localhost:5000/message")
	if err != nil {
		panic(err)
	}

	fmt.Println("---Got Message---")
	fmt.Printf("%+v\n", msg)
	fmt.Println("-----------------")
}

func GetMessage(url string) (ghtt.Message, error) {
	res, err := http.Get(url)
	if err != nil {
		return ghtt.Message{}, err
	}
	defer res.Body.Close()

	msgJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ghtt.Message{}, err
	}

	var msg ghtt.Message
	if err := json.Unmarshal(msgJson, &msg); err != nil {
		return ghtt.Message{}, err
	}

	return msg, nil
}
