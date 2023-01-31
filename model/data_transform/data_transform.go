package transform

import (
	"Go_cinema_reconstructed/model"
	pb "Go_cinema_reconstructed/proto"

	// u_pb "Go_cinema_reconstructed/user/proto"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

func TransformMovieRPC(movie *model.MovieRes) *pb.MovieInfoRes {

	updated_at := &google_protobuf.Timestamp{
		Seconds: movie.UpdatedAt.Unix(),
	}
	created_at := &google_protobuf.Timestamp{
		Seconds: movie.CreatedAt.Unix(),
	}
	director := &pb.Director{
		Firstname: movie.Director.Firstname,
		Lastname:  movie.Director.Lastname,
	}
	res := &pb.MovieInfoRes{
		Id:        movie.ID,
		Code:      movie.MovieID,
		Isbn:      movie.Isbn,
		Title:     movie.Title,
		Director:  director,
		UpdatedAt: updated_at,
		CreatedAt: created_at,
	}
	return res
}

func TransformMovieReqData(movie *pb.MovieInfoReq) *model.MovieReq {
	updated_at := time.Unix(movie.GetUpdatedAt().GetSeconds(), 0)
	created_at := time.Unix(movie.GetCreatedAt().GetSeconds(), 0)
	director := &model.Director{Firstname: movie.Director.Firstname, Lastname: movie.Director.Lastname}
	res := &model.MovieReq{
		MovieID:   movie.Code,
		Isbn:      movie.Isbn,
		Title:     movie.Title,
		Director:  *director,
		UpdatedAt: &updated_at,
		CreatedAt: &created_at,
	}
	return res
}

func TransformMovieResData(movie *pb.MovieInfoRes) *model.MovieRes {
	updated_at := time.Unix(movie.GetUpdatedAt().GetSeconds(), 0)
	created_at := time.Unix(movie.GetCreatedAt().GetSeconds(), 0)
	director := &model.Director{Firstname: movie.Director.Firstname, Lastname: movie.Director.Lastname}
	res := &model.MovieRes{
		ID:        movie.Id,
		MovieID:   movie.Code,
		Isbn:      movie.Isbn,
		Title:     movie.Title,
		Director:  *director,
		UpdatedAt: &updated_at,
		CreatedAt: &created_at,
	}
	return res
}

func TransformUserRPC(user *model.UserRes) *pb.UserInfoRes {
	if user == nil {
		return nil
	}

	updated_at := &google_protobuf.Timestamp{
		Seconds: user.UpdatedAt.Unix(),
	}
	created_at := &google_protobuf.Timestamp{
		Seconds: user.CreatedAt.Unix(),
	}
	res := &pb.UserInfoRes{
		Id:            user.ID,
		Code:          user.UserID,
		Email:         user.Email,
		Password:      user.Password,
		WatchedMovies: user.Watched_movies,
		UpdatedAt:     updated_at,
		CreatedAt:     created_at,
	}
	return res
}

func TransformUserReqData(user *pb.UserInfoReq) *model.UserReq {
	updated_at := time.Unix(user.GetUpdatedAt().GetSeconds(), 0)
	created_at := time.Unix(user.GetCreatedAt().GetSeconds(), 0)
	res := &model.UserReq{
		UserID:         user.Code,
		Email:          user.Email,
		Password:       user.Password,
		Watched_movies: user.WatchedMovies,
		UpdatedAt:      &updated_at,
		CreatedAt:      &created_at,
	}

	return res
}

func TransformUserResData(user *pb.UserInfoRes) *model.UserRes {
	updated_at := time.Unix(user.GetUpdatedAt().GetSeconds(), 0)
	created_at := time.Unix(user.GetCreatedAt().GetSeconds(), 0)
	res := &model.UserRes{
		ID:             user.Id,
		UserID:         user.Code,
		Email:          user.Email,
		Password:       user.Password,
		Watched_movies: user.WatchedMovies,
		UpdatedAt:      &updated_at,
		CreatedAt:      &created_at,
	}

	return res
}
