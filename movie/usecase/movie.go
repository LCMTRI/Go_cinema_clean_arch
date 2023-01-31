package usecase

import (
	"Go_cinema_clean_arch/model"
	"time"
	//"time"
)

// type MovieRepository interface {
// }

type movieUsecase struct {
	mRepo model.MovieRepository
}

func NewMovieUseCase(m model.MovieRepository) *movieUsecase {
	return &movieUsecase{
		mRepo: m,
	}
}

func (m *movieUsecase) GetAllMovies() ([]*model.MovieRes, error) {
	return m.mRepo.GetAll()
}

func (m *movieUsecase) GetMovieByID(movieID string) (*model.MovieRes, error) {
	return m.mRepo.GetByID(movieID)
}

func (m *movieUsecase) CreateMovie(movie *model.MovieReq) error {
	now := time.Now()
	movie.CreatedAt = &now
	movie.UpdatedAt = &now

	return m.mRepo.Create(movie)
}

func (m *movieUsecase) UpdateMovie(id string, movie *model.MovieReq) (int32, error) {
	now := time.Now()
	movie.UpdatedAt = &now

	return m.mRepo.Update(id, movie)
}

func (m *movieUsecase) DeleteMovie(id string) (int32, error) {
	return m.mRepo.Delete(id)
}

func (m *movieUsecase) GetWatchedMovies(idList []string) ([]*model.MovieRes, error) {
	return m.mRepo.GetWatchedMovies(idList)
}
