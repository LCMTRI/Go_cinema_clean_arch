// Dear programmer:
// When I wrote this code, only god and
// I knew how it worked.
// Now, only god knows it!
//
// Therefore, if you are trying to optimize
// this code and it fails (most surely),
// please increase this counter as a
// warning for the next person:
//
// total_hours_wasted_here = 39
//

package main

import (
	"Go_cinema_clean_arch/model"
	transform "Go_cinema_clean_arch/model/data_transform"
	pb "Go_cinema_clean_arch/proto"

	// u_pb "Go_cinema_clean_arch/user/proto"
	"context"
	"encoding/json"

	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	//"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPC Client
var clientUser pb.ComputeServiceClient
var clientMovie pb.ComputeServiceClient

// REST Server with gRPC functions inside handler functions
// for User
func getAllUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var Users []*pb.UserInfoRes

	req := &pb.Empty{}
	stream, err := clientUser.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", clientUser, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetUsers(_) = _, %v", clientUser, err)
		}
		Users = append(Users, row)
	}
	json.NewEncoder(c.Response().Writer).Encode(Users)
	return c.String(http.StatusOK, "")
}

func getUser(c echo.Context) error {
	userId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: userId}
	res, err := clientUser.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", clientUser, err)
	}

	if res != nil {
		json.NewEncoder(c.Response().Writer).Encode(res)
		return c.String(http.StatusOK, "")
	}

	return c.JSON(http.StatusNotFound, "Can't find the User with the given id")
}

func addUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = &pb.UserInfoReq{}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addUser: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	json.Unmarshal(b, &req)
	// c.Bind(req)
	res, err := clientUser.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateUser(_) = _, %v", clientUser, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateUser Id: %v", res.Value)
	} else {
		log.Printf("CreateUser Failed")
	}
	return c.String(http.StatusOK, "Successfully added a User")
}

func updateUser(c echo.Context) error {
	userId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.UserInfoReq{Code: userId}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
		return c.String(http.StatusBadRequest, "error reading the body")
	}

	json.Unmarshal(b, &req)
	res, err := clientUser.UpdateUser(ctx, &req)
	if err != nil {
		log.Fatalf("%v.UpdateUser(_) = _, %v", clientUser, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateUser Success")
		return c.String(http.StatusOK, "Successfully updated a User")
	} else {
		log.Printf("UpdateUser Failed")
		return c.String(http.StatusNotFound, "Can't find the User with the given id")
	}
}

func deleteUser(c echo.Context) error {
	userId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Id{Value: userId}
	res, err := clientUser.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteUser(_) = _, %v", clientUser, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteUser Success")
		return c.String(http.StatusOK, "Successfully deleted a User")
	} else {
		log.Printf("DeleteUser Failed")
		return c.String(http.StatusNotFound, "Can't find the User with the given id")
	}
}

// func getWatchedMovies(c echo.Context) error {
// 	userId := c.Param("id")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.Id{Value: userId}
// 	res, err := clientUser.GetWatchedMoviesGateway(ctx, req)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "Can't find the User with the given id")
// 	}
// 	json.NewEncoder(c.Response().Writer).Encode(res.WatchedMovies)
// 	return c.String(http.StatusOK, "")
// }

// for Movie
func getAllMovies(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var Movies []*model.MovieRes
	log.Println("we are here~")

	req := &pb.Empty{}
	stream, err := clientMovie.GetMovies(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovies(_) = _, %v", clientMovie, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMovies(_) = _, %v", clientMovie, err)
		}
		Movies = append(Movies, transform.TransformMovieResData(row))
	}
	json.NewEncoder(c.Response().Writer).Encode(Movies)
	return c.String(http.StatusOK, "")
}

func getMovie(c echo.Context) error {
	movieId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.Id{Value: movieId}
	res, err := clientMovie.GetMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovie(_) = _, %v", clientMovie, err)
	}

	if res != nil {
		json.NewEncoder(c.Response().Writer).Encode(res)
		return c.String(http.StatusOK, "")
	}

	return c.JSON(http.StatusNotFound, "Can't find the movie with the given id")
}

func addMovie(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.MovieInfoReq{}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addMovie: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	json.Unmarshal(b, &req)
	res, err := clientMovie.CreateMovie(ctx, &req)
	if err != nil {
		log.Fatalf("%v.CreateMovie(_) = _, %v", clientMovie, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateMovie Id: %v", res.Value)
	} else {
		log.Printf("CreateMovie Failed")
	}
	return c.String(http.StatusOK, "Successfully added a movie")
}

func updateMovie(c echo.Context) error {
	movieCode := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.MovieInfoReq{Code: movieCode}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
		return c.String(http.StatusBadRequest, "error reading the body")
	}

	json.Unmarshal(b, &req)
	res, err := clientMovie.UpdateMovie(ctx, &req)
	if err != nil {
		log.Fatalf("%v.UpdateMovie(_) = _, %v", clientMovie, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateMovie Success")
		return c.String(http.StatusOK, "Successfully updated a movie")
	} else {
		log.Printf("UpdateMovie Failed")
		return c.String(http.StatusNotFound, "Can't find the movie with the given id")
	}

}

func deleteMovie(c echo.Context) error {
	movieId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Id{Value: movieId}
	res, err := clientMovie.DeleteMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteMovie(_) = _, %v", clientMovie, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteMovie Success")
		return c.String(http.StatusOK, "Successfully deleted a movie")
	} else {
		log.Printf("DeleteMovie Failed")
		return c.String(http.StatusNotFound, "Can't find the movie with the given id")
	}
}

func main() {
	// User gRPC Client
	connUser, err1 := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err1 != nil {
		log.Fatalf("Did not connect: %v", err1)
	}
	defer connUser.Close()
	clientUser = pb.NewComputeServiceClient(connUser)
	fmt.Println("Welcom to User gRPC Server")

	// Movie gRPC Client
	connMovie, err2 := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err2 != nil {
		log.Fatalf("Did not connect: %v", err2)
	}
	defer connMovie.Close()
	clientMovie = pb.NewComputeServiceClient(connMovie)
	fmt.Println("Welcom to Movie gRPC Server")

	// REST Server
	e := echo.New()

	e.GET("/users", getAllUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", addUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	// e.GET("/users/:id/watched_movies", getWatchedMovies)

	e.GET("/movies", getAllMovies)
	e.GET("/movies/:id", getMovie)
	e.POST("/movies", addMovie)
	e.PUT("/movies/:id", updateMovie)
	e.DELETE("/movies/:id", deleteMovie)

	e.Start(":8000")
}
