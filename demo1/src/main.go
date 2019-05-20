package src

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ResponseData struct {
	Hostname string `json:"hostname"`
}

func main() {
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		_, _ = fmt.Fprint(w, ResponseData{Hostname: hostname})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
