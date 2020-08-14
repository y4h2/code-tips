package main

import (
	"eureka-go/service/fulfillment"
	"fmt"
	"log"
	"net/http"

	"github.com/hudl/fargo"
)

func main() {
	service := fulfillment.NewService()
	router := service.InitRoutes()

	http.Handle("/", router)
	port := 8082

	// register URL on eureka
	c := fargo.NewConn("http://127.0.0.1:8080/eureka/v2")
	i := fargo.Instance{
		HostName:       "i-6543",
		Port:           port,
		App:            "fulfillment",
		IPAddr:         "127.0.0.1",
		DataCenterInfo: fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:         fargo.UP,
	}

	c.RegisterInstance(&i)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
