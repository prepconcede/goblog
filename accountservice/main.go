package main

import (
	"fmt"
	"pivdoggo/goblog/accountservice/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
