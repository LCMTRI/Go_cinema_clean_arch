package repo

import (
	"Go_cinema_reconstructed/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type movieRepo struct {
	coll *mongo.Collection
}

func NewMovie(db *mongo.Database) *movieRepo {
	return &movieRepo{
		coll: db.Collection("Movies"),
	}
}

func (m *movieRepo) GetAll() ([]*model.MovieRes, error) {
	c, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	movies := make([]*model.MovieRes, 0)
	err = c.All(ctx, &movies)
	if err != nil {
		return nil, err
	}

	return movies, err
}

func (m *movieRepo) GetByID(movieID string) (*model.MovieRes, error) {
	resp := new(model.MovieRes)
	err := m.coll.FindOne(ctx, bson.M{"movie_id": movieID}).Decode(resp)
	if err != nil {
		log.Println(err)
	}
	return resp, err
}

func (m *movieRepo) Create(mInsert *model.MovieReq) error {
	_, err := m.coll.InsertOne(ctx, mInsert)
	// movie := &model.MovieRes{
	// 	ID:        add.InsertedID.(primitive.ObjectID).Hex(),
	// 	MovieID:   mInsert.MovieID,
	// 	Isbn:      mInsert.Isbn,
	// 	Title:     mInsert.Title,
	// 	Director:  mInsert.Director,
	// 	CreatedAt: mInsert.CreatedAt,
	// 	UpdatedAt: mInsert.UpdatedAt,
	// }
	return err
}

func (m *movieRepo) Update(movieID string, mUpdate *model.MovieReq) (int32, error) {
	// movie := &model.MovieReq{
	// 	MovieID:   mUpdate.MovieID,
	// 	Isbn:      mUpdate.Isbn,
	// 	Title:     mUpdate.Title,
	// 	Director:  mUpdate.Director,
	// 	CreatedAt: mUpdate.CreatedAt,
	// 	UpdatedAt: mUpdate.UpdatedAt,
	// }
	_, err := m.coll.UpdateOne(ctx, bson.M{"movie_id": movieID}, bson.D{{"$set", mUpdate}})
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (m *movieRepo) Delete(movieID string) (int32, error) {
	_, err := m.coll.DeleteOne(ctx, bson.M{"movie_id": movieID})
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (m *movieRepo) GetWatchedMovies(idList []string) ([]*model.MovieRes, error) {
	var res []*model.MovieRes
	movies, _ := m.GetAll()
	for _, movieId := range idList {
		for _, movie := range movies {
			if movieId == movie.MovieID {
				res = append(res, movie)
			}
		}
	}
	return res, nil
}
