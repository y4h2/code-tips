package main

import (
	"eureka-go/service/catalog"
	"fmt"
	"log"
	"net/http"

	"github.com/hudl/fargo"
)

func main() {

	c := fargo.NewConn("http://127.0.0.1:8080/eureka/v2")
	app, err := c.GetApp("fulfillment")
	if err != nil {
		log.Fatalf("fail to connect eureka: %v", err)
	}

	webClient := catalog.NewFulfillmentWebClient(fmt.Sprintf("http://%s:%d/skus", app.Instances[0].IPAddr, app.Instances[0].Port))

	service := catalog.NewService(webClient)
	router := service.InitRoutes()

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
