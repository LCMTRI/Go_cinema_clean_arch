package model

import "time"

// User
type UserRes struct {
	ID             string     `json:"_id" bson:"_id"`
	UserID         string     `json:"user_id" bson:"user_id"`
	Email          string     `json:"email" bson:"email"`
	Password       string     `json:"password" bson:"password"`
	Watched_movies []string   `json:"watched_movies" bson:"watched_movies"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedAt      *time.Time `json:"created_at"`
}

// User model for creation
type UserReq struct {
	UserID         string     `json:"user_id" bson:"user_id"`
	Email          string     `json:"email" bson:"email"`
	Password       string     `json:"password" bson:"password"`
	Watched_movies []string   `json:"watched_movies" bson:"watched_movies"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedAt      *time.Time `json:"created_at"`
}

// UserUsecase represent the user's usecases
type UserUseCase interface {
	GetAllUsers() ([]*UserRes, error)
	GetUserByID(id string) (*UserRes, error)
	CreateUser(user *UserReq) error
	UpdateUser(id string, user *UserReq) (int32, error)
	DeleteUser(id string) (int32, error)
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	GetAll() ([]*UserRes, error)
	GetByID(id string) (*UserRes, error)
	Create(user *UserReq) error
	Update(userID string, user *UserReq) (int32, error)
	Delete(userID string) (int32, error)
}
