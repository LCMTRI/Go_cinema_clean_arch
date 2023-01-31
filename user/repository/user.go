package repo

import (
	"Go_cinema_clean_arch/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type userRepo struct {
	coll *mongo.Collection
}

func NewUser(db *mongo.Database) *userRepo {
	return &userRepo{
		coll: db.Collection("Users"),
	}
}

func (u *userRepo) GetAll() ([]*model.UserRes, error) {
	c, err := u.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	users := make([]*model.UserRes, 0)
	err = c.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (u *userRepo) GetByID(userID string) (*model.UserRes, error) {
	resp := new(model.UserRes)
	err := u.coll.FindOne(ctx, bson.M{"user_id": userID}).Decode(resp)
	return resp, err
}

func (u *userRepo) Create(uInsert *model.UserReq) error {
	// user := &model.UserReq{
	// 	UserID:         uInsert.UserID,
	// 	Email:          uInsert.Email,
	// 	Password:       uInsert.Password,
	// 	Watched_movies: uInsert.Watched_movies,
	// 	CreatedAt:      uInsert.CreatedAt,
	// 	UpdatedAt:      uInsert.UpdatedAt,
	// }
	_, err := u.coll.InsertOne(ctx, uInsert)
	return err
}

func (u *userRepo) Update(userID string, uUpdate *model.UserReq) (int32, error) {
	// user := &model.UserReq{
	// 	UserID:         uUpdate.UserID,
	// 	Email:          uUpdate.Email,
	// 	Password:       uUpdate.Password,
	// 	Watched_movies: uUpdate.Watched_movies,
	// 	CreatedAt:      uUpdate.CreatedAt,
	// 	UpdatedAt:      uUpdate.UpdatedAt,
	// }
	_, err := u.coll.UpdateOne(ctx, bson.M{"user_id": userID}, bson.D{{"$set", uUpdate}})
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (u *userRepo) Delete(userID string) (int32, error) {
	_, err := u.coll.DeleteOne(ctx, bson.M{"user_id": userID})
	if err != nil {
		return 0, err
	}
	return 1, nil
}
