package main

import (
	"flag"
	"fmt"
	"service/server"
	"service/utils"
	"service/utils/postgresClient"
)

func main() {
	// parse flag
	flag.Parse()

	// load .env
	success := utils.InitEnv()
	if !success {
		return
	}

	// init DB
	success = postgresClient.InitDB()
	if !success {
		panic("InitDB: fail")
	} else {
		fmt.Println("InitDB: success")
	}

	// start gin server
	server.InitGinServer()
}
