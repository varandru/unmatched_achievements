package main

import (
	"fmt"

	"./http_server"
)

func main() {
	var handler http_server.HttpUrlsHandler

	err := handler.ListenAndServe()
	fmt.Println(err)
}
