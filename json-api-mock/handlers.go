package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to API server written in golang!")
}

func API(w http.ResponseWriter, r *http.Request) {
	responses := Responses{
		Response{Name: "John"},
		Response{Name: "Bob"},
		Response{Name: "Shelly"},
	}

	if err := json.NewEncoder(w).Encode(responses); err != nil {
		panic(err)
	}
}
