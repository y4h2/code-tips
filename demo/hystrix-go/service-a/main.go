package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {

	// override some default values for the Hystrix breaker
	hystrix.DefaultVolumeThreshold = 3
	hystrix.DefaultErrorPercentThreshold = 75
	hystrix.DefaultTimeout = 500
	hystrix.DefaultSleepWindow = 3500

	// export Hystrix stream
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "8081"), hystrixStreamHandler)

	http.HandleFunc("/", logger(HandleSubsystem))

	http.ListenAndServe(":8080", nil)
}

// HandleSubsystem send request to sub-system and extracts its response
func HandleSubsystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resultCh := make(chan []byte)
	errCh := hystrix.Go("my_command", func() error {
		resp, err := http.Get("http://localhost:9090")
		if err != nil {
			return err
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		resultCh <- b
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		log.Println("success to get response from sub-system:", string(res))
		w.WriteHeader(http.StatusOK)
	case err := <-errCh:
		log.Println("failed to get response from sub-system:", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

// log is Handler wrapper function for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}
