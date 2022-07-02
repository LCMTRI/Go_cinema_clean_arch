package handler

import (
	"Go_cinema_reconstructed/model"
	transform "Go_cinema_reconstructed/model/data_transform"
	pb "Go_cinema_reconstructed/proto"
	"context"
	"log"

	// "time"

	//"net/http"
	"strconv"

	//"github.com/labstack/echo/v4"
	// google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

type MovieHandler struct {
	pb.UnimplementedComputeServiceServer
	m model.MovieUseCase
}

func NewMovieHandler(m model.MovieUseCase) *MovieHandler {
	return &MovieHandler{m: m}
}

func NewMovieServerGrpc(gserver *grpc.Server, movieUcase model.MovieUseCase) {

	movieServer := &MovieHandler{
		m: movieUcase,
	}

	pb.RegisterComputeServiceServer(gserver, movieServer)
}

// func (h *MovieHandler) GetMovies(c echo.Context) error {
// 	resp, err := h.m.GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }

// ============== gRPC server ===============
// ==========================================

func (h *MovieHandler) GetMovies(in *pb.Empty, stream pb.ComputeService_GetMoviesServer) error {
	movies, _ := h.m.GetAllMovies()
	for _, movie := range movies {
		if err := stream.Send(transform.TransformMovieRPC(movie)); err != nil {
			return err
		}
	}

	return nil
}

func (h *MovieHandler) GetMovie(ctx context.Context, in *pb.Id) (*pb.MovieInfoRes, error) {
	log.Printf("Received: %v", in)

	movie, err := h.m.GetMovieByID(in.Value)
	if err != nil {
		log.Println("err here!")
		return nil, err
	}

	return transform.TransformMovieRPC(movie), nil
}

func (h *MovieHandler) CreateMovie(ctx context.Context, in *pb.MovieInfoReq) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := &pb.Id{}

	mInsert := transform.TransformMovieReqData(in)
	mInsert.MovieID = strconv.Itoa(model.IdCounter)
	model.IdCounter++
	err := h.m.CreateMovie(mInsert)
	if err != nil {
		return nil, err
	}
	res.Value = mInsert.MovieID

	return res, nil
}

func (h *MovieHandler) UpdateMovie(ctx context.Context, in *pb.MovieInfoReq) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	mUpdate := transform.TransformMovieReqData(in)

	resp, err := h.m.UpdateMovie(mUpdate.MovieID, mUpdate)
	if err != nil {
		return nil, err
	}
	res.Value = resp

	return &res, nil
}

func (h *MovieHandler) DeleteMovie(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}

	resp, err := h.m.DeleteMovie(in.Value)
	if err != nil {
		return nil, err
	}
	res.Value = resp

	return &res, nil
}

func (h *MovieHandler) GetWatchedMoviesUser(in *pb.MovieList, stream pb.ComputeService_GetWatchedMoviesUserServer) error {
	movies, _ := h.m.GetWatchedMovies(in.WatchedMovies)
	for _, movie := range movies {
		if err := stream.Send(transform.TransformMovieRPC(movie)); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
