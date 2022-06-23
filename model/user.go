package model

// User
type User struct {
	ID             string   `json:"_id" bson:"_id"`
	UserID         string   `json:"userID" bson:"userID"`
	Email          string   `json:"email" bson:"email"`
	Password       string   `json:"password" bson:"password"`
	Watched_movies []string `json:"watched_movies" bson:"watched_movies"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
}

// UserRepository represent the user's repository contract
type UserRepository interface {
}
