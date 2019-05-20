package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ResponseData struct {
	Hostname string `json:"hostName"`
	Message  string `json:"message"`
}

func main() {
	// api interface
	http.HandleFunc("/api/k8s", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		resp := ResponseData{
			Hostname: hostname,
			Message:  "hello k8s "}
		data, _ := json.Marshal(resp)
		_, _ = fmt.Fprint(w, string(data))
	})

	// frontend route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
