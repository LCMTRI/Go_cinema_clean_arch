package main

import (
	movie_handler "Go_cinema_clean_arch/movie/handler"
	movie_repo "Go_cinema_clean_arch/movie/repository"
	movie_usecase "Go_cinema_clean_arch/movie/usecase"
	"context"

	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var MovieHandler *movie_handler.MovieHandler

func main() {
	var URI = "mongodb+srv://lecaominhtri0701:lecaominhtri@cluster0.x7apzya.mongodb.net/?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}
	movieDatabase := client.Database("Movie_service_database")

	movie := movie_repo.NewMovie(movieDatabase)
	movieUseCase := movie_usecase.NewMovieUseCase(movie)

	// gRPC Server port :8001
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	movie_handler.NewMovieServerGrpc(server, movieUseCase)

	log.Printf("Movie gRPC Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	//MovieHandler = movie_handler.NewMovieHandler(movieUseCase)
}
