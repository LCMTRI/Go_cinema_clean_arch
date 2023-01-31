package gw_handler

import (
	//m_handler "Go_cinema_reconstructed/movie/handler"
	//u_handler "Go_cinema_reconstructed/user/handler"

	"Go_cinema_reconstructed/model"
	transform "Go_cinema_reconstructed/model/data_transform"

	//u_handler "Go_cinema_reconstructed/user/handler"
	pb "Go_cinema_reconstructed/proto"
	//u_pb "Go_cinema_reconstructed/user/proto"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
//clientUser  u_pb.UserServiceClient
//ClientMovie m_pb.MovieServiceClient
)

// func SetClientMovie(c m_pb.MovieServiceClient) {
// 	clientMovie = c
// }

// type Handler interface {
// 	GetMovies(c echo.Context) error
// 	GetMovie(c echo.Context) error
// 	CreateMovie(c echo.Context) error
// 	UpdateMovie(c echo.Context) error
// 	DeleteMovie(c echo.Context) error

// 	GetUsers(c echo.Context) error
// 	GetUser(c echo.Context) error
// 	CreateUser(c echo.Context) error
// 	UpdateUser(c echo.Context) error
// 	DeleteUser(c echo.Context) error
// }

type GatewayHandler struct {
	//Handler
	clientMovie pb.ComputeServiceClient
	clientUser  pb.ComputeServiceClient
}

func NewGatewayHandler(clientM pb.ComputeServiceClient, clientU pb.ComputeServiceClient) *GatewayHandler {
	return &GatewayHandler{
		clientMovie: clientM,
		clientUser:  clientU,
	}
}

// Movie Handler
func (h *GatewayHandler) GetMovies(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var movies []*model.MovieRes

	req := &pb.Empty{}
	stream, err := h.clientMovie.GetMovies(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovies(_) = _, %v", h.clientMovie, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetMovies(_) = _, %v", h.clientMovie, err)
		}
		movies = append(movies, transform.TransformMovieResData(row))
	}
	json.NewEncoder(c.Response().Writer).Encode(movies)
	return c.String(http.StatusOK, "")
}

func (h *GatewayHandler) GetMovie(c echo.Context) error {
	movieId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.Id{Value: movieId}
	res, err := h.clientMovie.GetMovie(ctx, req)
	if err != nil {
		log.Printf("%v.GetMovies(_) = _, %v", h.clientMovie, err)
	}

	if res != nil {
		json.NewEncoder(c.Response().Writer).Encode(res)
		return c.String(http.StatusOK, "")
	}

	return c.JSON(http.StatusNotFound, "Can't find the movie with the given id")
}

func (h *GatewayHandler) AddMovie(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.MovieInfoReq{}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addMovie: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	json.Unmarshal(b, &req)
	res, err := h.clientMovie.CreateMovie(ctx, &req)
	if err != nil {
		log.Fatalf("%v.CreateMovie(_) = _, %v", h.clientMovie, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateMovie Id: %v", res.Value)
	} else {
		log.Printf("CreateMovie Failed")
	}
	return c.String(http.StatusOK, "Successfully added a Movie")
}

func (h *GatewayHandler) UpdateMovie(c echo.Context) error {
	movieCode := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.MovieInfoReq{Code: movieCode}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
		return c.String(http.StatusBadRequest, "error reading the body")
	}

	json.Unmarshal(b, &req)
	res, err := h.clientMovie.UpdateMovie(ctx, &req)
	if err != nil {
		log.Fatalf("%v.UpdateMovie(_) = _, %v", h.clientMovie, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateMovie Success")
		return c.String(http.StatusOK, "Successfully updated a movie")
	} else {
		log.Printf("UpdateMovie Failed")
		return c.String(http.StatusNotFound, "Can't find the movie with the given id")
	}

}

func (h *GatewayHandler) DeleteMovie(c echo.Context) error {
	movieId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.Id{Value: movieId}
	res, err := h.clientMovie.DeleteMovie(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteMovie(_) = _, %v", h.clientMovie, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteMovie Success")
		return c.String(http.StatusOK, "Successfully deleted a movie")
	} else {
		log.Printf("DeleteMovie Failed")
		return c.String(http.StatusNotFound, "Can't find the movie with the given id")
	}
}

// User Handler
func (h *GatewayHandler) GetUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users []*model.UserRes

	req := &pb.Empty{}
	stream, err := h.clientUser.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", h.clientUser, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetUsers(_) = _, %v", h.clientUser, err)
		}
		users = append(users, transform.TransformUserResData(row))
	}
	json.NewEncoder(c.Response().Writer).Encode(users)
	return c.String(http.StatusOK, "")
}

func (h *GatewayHandler) GetUser(c echo.Context) error {
	userId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.Id{Value: userId}
	res, err := h.clientUser.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", h.clientUser, err)
	}

	if res != nil {
		json.NewEncoder(c.Response().Writer).Encode(res)
		return c.String(http.StatusOK, "")
	}

	return c.JSON(http.StatusNotFound, "Can't find the User with the given id")
}

func (h *GatewayHandler) AddUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.UserInfoReq{}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addUser: %s", err)
		return c.String(http.StatusBadRequest, "")
	}

	json.Unmarshal(b, &req)
	// c.Bind(req)
	res, err := h.clientUser.CreateUser(ctx, &req)
	if err != nil {
		log.Fatalf("%v.CreateUser(_) = _, %v", h.clientUser, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateUser Id: %v", res.Value)
	} else {
		log.Printf("CreateUser Failed")
	}
	return c.String(http.StatusOK, "Successfully added a User")
}

func (h *GatewayHandler) UpdateUser(c echo.Context) error {
	userId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer c.Request().Body.Close()
	var req = pb.UserInfoReq{Code: userId}

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatalln(err)
		return c.String(http.StatusBadRequest, "error reading the body")
	}

	json.Unmarshal(b, &req)
	res, err := h.clientUser.UpdateUser(ctx, &req)
	if err != nil {
		log.Fatalf("%v.UpdateUser(_) = _, %v", h.clientUser, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateUser Success")
		return c.String(http.StatusOK, "Successfully updated a User")
	} else {
		log.Printf("UpdateUser Failed")
		return c.String(http.StatusNotFound, "Can't find the User with the given id")
	}
}

func (h *GatewayHandler) DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.Id{Value: userId}
	res, err := h.clientUser.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteUser(_) = _, %v", h.clientUser, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteUser Success")
		return c.String(http.StatusOK, "Successfully deleted a User")
	} else {
		log.Printf("DeleteUser Failed")
		return c.String(http.StatusNotFound, "Can't find the User with the given id")
	}
}

func (h *GatewayHandler) GetWatchedMovies(c echo.Context) error {
	userId := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &pb.Id{Value: userId}
	res, err := h.clientUser.GetWatchedMoviesGateway(ctx, req)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Can't find the User with the given id")
	}
	var watched_movies []*model.MovieRes
	for _, movie := range res.WatchedMovies {
		watched_movies = append(watched_movies, transform.TransformMovieResData(movie))
	}
	json.NewEncoder(c.Response().Writer).Encode(watched_movies)
	return c.String(http.StatusOK, "")
}
