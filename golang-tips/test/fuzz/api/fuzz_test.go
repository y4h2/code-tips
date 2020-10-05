package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func TestFuzz(t *testing.T) {
	Routes()

	f := fuzz.New()
	var data []byte

	for i := 0; i < 1000; i++ {
		f.Fuzz(&data)
		r, _ := http.NewRequest("POST", "/process", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
	}

}
