package main

import (
	"fmt"

	"github.com/hudl/fargo"
)

func main() {
	// For a real app, you'd bind a user-provided service with eureka
	// credentials and URL.
	c := fargo.NewConn("http://127.0.0.1:8080/eureka/v2")
	i := fargo.Instance{
		HostName:         "i-6543",
		Port:             9090,
		App:              "TESTAPP",
		IPAddr:           "127.0.0.10",
		VipAddress:       "127.0.0.10",
		SecureVipAddress: "127.0.0.10",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}

	c.RegisterInstance(&i)

	f, _ := c.GetApps()

	for key, theApp := range f {
		fmt.Println("App:", key, " First Host Name:", theApp.Instances[0].HostName)
	}

	app, _ := c.GetApp("TESTAPP")

	fmt.Printf("%v\n", app.Instances[0].IPAddr)
	fmt.Printf("%v\n", app.Instances[0].Port)
}
