package main

import (
	"service/test/service"
	"service/utils"
	// "service/utils/postgresClient"
)

func main() {
	utils.InitEnv()
	service.EventCount()
}
