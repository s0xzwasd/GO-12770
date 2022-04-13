package main

import (
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	isTracing := grpc.EnableTracing
	fmt.Println(isTracing)
}
