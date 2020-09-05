package main

import (
	"net/http"

	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func test1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8000", handler)
}

func test2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	http.ListenAndServe(":8000", handler)
}
