package main

import (
	"fmt"
	"server/src/http_server"
)

func main() {
	var handler http_server.HttpUrlsHandler
	handler.Initialize()

	err := handler.Start()
	fmt.Println(err)
}
