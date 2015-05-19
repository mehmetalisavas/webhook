package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", payload)
	http.ListenAndServe(":4567", nil)
}

func payload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, My name is Mehmet Ali ", r.URL.Path[1:])
	var pay map[string]interface{}

	json.NewDecoder(r.Body).Decode(&pay)
	fmt.Fprintf(w, pay, r.URL.Path[1:])
}
