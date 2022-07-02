package handler

import (
	"Go_cinema_reconstructed/model"
	transform "Go_cinema_reconstructed/model/data_transform"
	pb "Go_cinema_reconstructed/proto"
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

type UserHandler struct {
	pb.UnimplementedComputeServiceServer
	u           model.UserUseCase
	clientMovie pb.ComputeServiceClient
}

// func NewUserHandler(u model.UserUseCase, clientMovie pb.ComputeServiceClient) *UserHandler {
// 	return &UserHandler{
// 		u:           u,
// 		clientMovie: clientMovie,
// 	}
// }

func NewUserServerGrpc(gserver *grpc.Server, userUsecase model.UserUseCase, clientM pb.ComputeServiceClient) {

	userServer := &UserHandler{
		u:           userUsecase,
		clientMovie: clientM,
	}

	pb.RegisterComputeServiceServer(gserver, userServer)
}

func (h *UserHandler) GetUsers(in *pb.Empty, stream pb.ComputeService_GetUsersServer) error {
	users, _ := h.u.GetAllUsers()
	for _, user := range users {
		if err := stream.Send(transform.TransformUserRPC(user)); err != nil {
			return err
		}
	}

	return nil
}

func (h *UserHandler) GetUser(ctx context.Context, in *pb.Id) (*pb.UserInfoRes, error) {
	log.Printf("Received: %v", in)

	user, err := h.u.GetUserByID(in.Value)
	if err != nil {
		return nil, err
	}

	return transform.TransformUserRPC(user), nil
}

func (h *UserHandler) CreateUser(ctx context.Context, in *pb.UserInfoReq) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := &pb.Id{}

	mInsert := transform.TransformUserReqData(in)
	mInsert.UserID = strconv.Itoa(model.IdCounter)
	model.IdCounter++
	err := h.u.CreateUser(mInsert)
	if err != nil {
		return nil, err
	}
	res.Value = mInsert.UserID

	return res, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, in *pb.UserInfoReq) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}
	mUpdate := transform.TransformUserReqData(in)

	resp, err := h.u.UpdateUser(mUpdate.UserID, mUpdate)
	if err != nil {
		return nil, err
	}
	res.Value = resp

	return &res, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := pb.Status{}

	resp, err := h.u.DeleteUser(in.Value)
	if err != nil {
		return nil, err
	}
	res.Value = resp

	return &res, nil
}

func (h *UserHandler) GetWatchedMoviesGateway(ctx context.Context, in *pb.Id) (*pb.MovieInfoList, error) {
	log.Printf("Received user id: %v", in)
	req := &pb.MovieList{}
	//var watched_movies []string
	user, err := h.GetUser(ctx, in)
	if err != nil {
		log.Fatalln(err)
	}

	req.WatchedMovies = user.WatchedMovies
	//ctx := context.Background()
	stream, err := h.clientMovie.GetWatchedMoviesUser(ctx, req)
	var movies []*pb.MovieInfoRes
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
		movies = append(movies, row)

	}
	res := pb.MovieInfoList{WatchedMovies: movies}
	return &res, nil
}
