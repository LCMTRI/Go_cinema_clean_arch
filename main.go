package main

import (
	gw_handler "Go_cinema_reconstructed/gateway/handler"
	"Go_cinema_reconstructed/gateway/route"
	pb "Go_cinema_reconstructed/proto"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// User gRPC Client
	connUser, err1 := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("Did not connect: %v", err1)
	}
	defer connUser.Close()
	clientUser := pb.NewComputeServiceClient(connUser)
	fmt.Println("Welcom to User gRPC Server")

	// Movie gRPC Client
	connMovie, err2 := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err2 != nil {
		log.Fatalf("Did not connect: %v", err2)
	}
	defer connMovie.Close()
	clientMovie := pb.NewComputeServiceClient(connMovie)
	fmt.Println("Welcom to Movie gRPC Server")

	e := echo.New()
	handler := gw_handler.NewGatewayHandler(clientMovie, clientUser)

	route.Private(e, handler, connMovie)

	e.Start(":8000")
}
