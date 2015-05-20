package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", payload)
	http.HandleFunc("/repos/mehmetalisavas/webhook/hooks", List)
	http.ListenAndServe(":4567", nil)
}
// payload is created for github webhook
func payload(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi, My name is Mehmet Ali ", r.URL.Path[1:])
	var pay map[string]interface{}

	// json.Unmarshal(r, &pay)
	json.NewDecoder(r.Body).Decode(&pay)
	// fmt.Fprintf(w, pay, r.URL.Path[1:])
	fmt.Println(pay)
	// fmt.Println(w, pay, r.URL.Path[1:])
	fmt.Println("SomethingLikeThis")
}

func List(w http.ResponseWriter, r *http.Request) {
	var pay map[string]interface{}

	// json.Unmarshal(r, &pay)
	json.NewDecoder(r.Body).Decode(&pay)
	// fmt.Fprintf(w, pay, r.URL.Path[1:])
	fmt.Println(pay)
	// fmt.Println(w, pay, r.URL.Path[1:])
	fmt.Println("SomethingLikeThis")
}
