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

func test(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	resp := ResponseData{}
	resp.Hostname = hostname
	_, _ = fmt.Fprint(w, resp)
}

func main() {
	http.HandleFunc("/api/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
