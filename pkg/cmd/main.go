package main

import (
	"fmt"
	grpcserver "grpc-server/pkg/server"
)

func main() {
	fmt.Print("server started !!!\n")
	grpcserver.StartServer()
	fmt.Print("server started !!!\n")
}
