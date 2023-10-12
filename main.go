package main

import "lite-gateway/gateway"

func main() {
	gw := gateway.NewGateway()

	if err := gw.Run(":8080"); err != nil {
		panic(err)
	}
}
