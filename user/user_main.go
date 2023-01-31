package main

import (
	pb "Go_cinema_clean_arch/proto"
	user_handler "Go_cinema_clean_arch/user/handler"
	user_repo "Go_cinema_clean_arch/user/repository"
	user_usecase "Go_cinema_clean_arch/user/usecase"
	"context"
	"fmt"

	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserHandler *user_handler.UserHandler

func main() {
	// gRPC Client to port :8001
	connMovie, err2 := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err2 != nil {
		log.Fatalf("Did not connect: %v", err2)
	}
	defer connMovie.Close()
	clientMovie := pb.NewComputeServiceClient(connMovie)
	fmt.Println("Welcom to Movie gRPC Server")

	// Database connection
	var URI = "mongodb+srv://lecaominhtri0701:lecaominhtri@cluster0.x7apzya.mongodb.net/?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}
	userDatabase := client.Database("User_service_database")

	user := user_repo.NewUser(userDatabase)
	userUseCase := user_usecase.NewUserUseCase(user)

	// gRPC Server port :8002
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	user_handler.NewUserServerGrpc(server, userUseCase, clientMovie)

	log.Printf("User gRPC Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
