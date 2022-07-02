package model

import "time"

var IdCounter = 1

// Movie
type MovieRes struct {
	ID        string     `json:"_id" bson:"_id"`
	MovieID   string     `json:"movie_id" bson:"movie_id"`
	Isbn      string     `json:"isbn" bson:"isbn"`
	Title     string     `json:"title" bson:"title"`
	Director  Director   `json:"director" bson:"director"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

type Director struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

// Movie model for creation
type MovieReq struct {
	MovieID   string     `json:"movie_id" bson:"movie_id"`
	Isbn      string     `json:"isbn" bson:"isbn"`
	Title     string     `json:"title" bson:"title"`
	Director  Director   `json:"director" bson:"director"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// MovieUsecase represent the movie's usecases
type MovieUseCase interface {
	GetAllMovies() ([]*MovieRes, error)
	GetMovieByID(id string) (*MovieRes, error)
	CreateMovie(movie *MovieReq) error
	UpdateMovie(id string, movie *MovieReq) (int32, error)
	DeleteMovie(id string) (int32, error)
	GetWatchedMovies(idList []string) ([]*MovieRes, error)
}

// MovieRepository represent the movie's repository contract
type MovieRepository interface {
	GetAll() ([]*MovieRes, error)
	GetByID(id string) (*MovieRes, error)
	Create(movie *MovieReq) error
	Update(movieID string, movie *MovieReq) (int32, error)
	Delete(movieID string) (int32, error)
	GetWatchedMovies(idList []string) ([]*MovieRes, error)
}
