package main

import (
	"fmt"
	"simple-api-server/cmd/server"
)

const port = 8080

func main() {
	fmt.Println("Simple API Server")
	server.Main(port)
}
