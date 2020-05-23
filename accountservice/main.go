package main

import (
	"fmt"
	"pivdoggo/goblog/accountservice/dbclient"
	"pivdoggo/goblog/accountservice/services"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initBoltClient()
	service.StartWebServer("6767")
}

func initBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
