package usecase

import (
	"Go_cinema_reconstructed/model"
	"time"
)

type userUsecase struct {
	uRepo model.UserRepository
}

func NewUserUseCase(u model.UserRepository) *userUsecase {
	return &userUsecase{
		uRepo: u,
	}
}

func (u *userUsecase) GetAllUsers() ([]*model.UserRes, error) {
	return u.uRepo.GetAll()
}

func (u *userUsecase) GetUserByID(userID string) (*model.UserRes, error) {
	return u.uRepo.GetByID(userID)
}

func (u *userUsecase) CreateUser(user *model.UserReq) error {
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now

	return u.uRepo.Create(user)
}

func (u *userUsecase) UpdateUser(id string, user *model.UserReq) (int32, error) {
	now := time.Now()
	user.UpdatedAt = &now

	return u.uRepo.Update(id, user)
}

func (u *userUsecase) DeleteUser(id string) (int32, error) {
	return u.uRepo.Delete(id)
}
