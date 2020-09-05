package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// rand.Seed()

	http.HandleFunc("/", logger(HandleHeavyJob))

	fmt.Println("==> Sub server is started")
	log.Println("listening on :9090")
	http.ListenAndServe(":9090", nil)
}

// HandleHeavyJob send request to sub-system and extracts its response
func HandleHeavyJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// This takes time and sometimes fail...

	// n := rand.Intn(1000)
	n := 1000
	log.Printf("delay %d", n)
	time.Sleep(time.Duration(n) * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// log is Handler wrapper function for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}
