package model

// Movie
type Movie struct {
	ID       string   `json:"_id" bson:"_id"`
	MovieID  string   `json:"movieID" bson:"movieID"`
	Isbn     string   `json:"isbn" bson:"isbn"`
	Title    string   `json:"title" bson:"title"`
	Director Director `json:"director" bson:"director"`
}

type Director struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

// MovieUsecase represent the movie's usecases
type MovieUsecase interface {
}

// MovieRepository represent the movie's repository contract
type MovieRepository interface {
}
