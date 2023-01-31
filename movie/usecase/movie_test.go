package usecase

import (
	"testing"

	"Go_cinema_reconstructed/model"
	mock_model "Go_cinema_reconstructed/model/mock"

	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetAllMovie(t *testing.T) {
	now := time.Now()
	movies := []*model.MovieRes{
		{
			ID:      "20221",
			MovieID: "1",
			Isbn:    "111",
			Title:   "Toy Story",
			Director: model.Director{
				Firstname: "John",
				Lastname:  "Lasseter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20222",
			MovieID: "2",
			Isbn:    "112",
			Title:   "A Bug's Life",
			Director: model.Director{
				Firstname: "John",
				Lastname:  "Lasseter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20223",
			MovieID: "3",
			Isbn:    "113",
			Title:   "Toy Story 2",
			Director: model.Director{
				Firstname: "John",
				Lastname:  "Lasseter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20224",
			MovieID: "4",
			Isbn:    "114",
			Title:   "Monsters, Inc.",
			Director: model.Director{
				Firstname: "Pete",
				Lastname:  "Docter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20225",
			MovieID: "5",
			Isbn:    "115",
			Title:   "Finding Nemo",
			Director: model.Director{
				Firstname: "Andrew",
				Lastname:  "Stanton",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20226",
			MovieID: "6",
			Isbn:    "116",
			Title:   "The Incredibles",
			Director: model.Director{
				Firstname: "Brad",
				Lastname:  "Bird",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20227",
			MovieID: "7",
			Isbn:    "117",
			Title:   "Cars",
			Director: model.Director{
				Firstname: "John",
				Lastname:  "Lasseter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20228",
			MovieID: "8",
			Isbn:    "118",
			Title:   "Ratatouille",
			Director: model.Director{
				Firstname: "Brad",
				Lastname:  "Bird",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "20229",
			MovieID: "9",
			Isbn:    "119",
			Title:   "WALL-E",
			Director: model.Director{
				Firstname: "Andrew",
				Lastname:  "Stanton",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "202210",
			MovieID: "10",
			Isbn:    "120",
			Title:   "Up",
			Director: model.Director{
				Firstname: "Pete",
				Lastname:  "Docter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "202211",
			MovieID: "11",
			Isbn:    "121",
			Title:   "Toy Story 3",
			Director: model.Director{
				Firstname: "Lee",
				Lastname:  "Unkrich",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			ID:      "202212",
			MovieID: "12",
			Isbn:    "122",
			Title:   "Cars 2",
			Director: model.Director{
				Firstname: "John",
				Lastname:  "Lasseter",
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}
	testCases := []struct {
		name       string
		buildStubs func(mock_movie *mock_model.MockMovieRepository)
	}{
		{
			name: "OK",
			buildStubs: func(mock_movie *mock_model.MockMovieRepository) {
				mock_movie.EXPECT().GetAll().Times(1).Return(movies, nil)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_model.NewMockMovieRepository(ctrl)
			u := NewMovieUseCase(repo)
			tc.buildStubs(repo)

			list, err := u.GetAllMovies()
			require.NoError(t, err)
			require.NotEmpty(t, list)
			require.Equal(t, list, movies)
			require.ElementsMatch(t, list, movies)

		})
	}
}

func TestGetMovieByID(t *testing.T) {
	now := time.Now()
	movie := &model.MovieRes{
		ID:      "20227",
		MovieID: "7",
		Isbn:    "117",
		Title:   "Cars",
		Director: model.Director{
			Firstname: "John",
			Lastname:  "Lasseter",
		},
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	testCases := []struct {
		name          string
		buildStubs    func(mock_movie *mock_model.MockMovieRepository)
		checkResponse func(movie *model.MovieRes, m *model.MovieRes, err error)
	}{
		{
			name: "OK",
			buildStubs: func(mock_movie *mock_model.MockMovieRepository) {
				mock_movie.EXPECT().GetByID(gomock.Eq(movie.MovieID)).Times(1).Return(movie, nil)
			},
			checkResponse: func(movie *model.MovieRes, m *model.MovieRes, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, m)
				require.Equal(t, m.ID, movie.ID)
				require.Equal(t, m.MovieID, movie.MovieID)
				require.Equal(t, m.Isbn, movie.Isbn)
				require.Equal(t, m.Title, movie.Title)
				require.Equal(t, m.Director, movie.Director)
				require.WithinDuration(t, *m.CreatedAt, *movie.CreatedAt, time.Millisecond)
				require.WithinDuration(t, *m.UpdatedAt, *movie.UpdatedAt, time.Millisecond)
			},
		},
		{
			name: "Notfound",
			buildStubs: func(mock_movie *mock_model.MockMovieRepository) {
				mock_movie.EXPECT().GetByID(gomock.Eq(movie.MovieID)).Times(1).Return(&model.MovieRes{}, mongo.ErrNoDocuments)
			},
			checkResponse: func(movie *model.MovieRes, m *model.MovieRes, err error) {
				require.Error(t, err)
				require.EqualError(t, err, mongo.ErrNoDocuments.Error())
				require.Equal(t, &model.MovieRes{}, m)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_model.NewMockMovieRepository(ctrl)
			u := NewMovieUseCase(repo)
			tc.buildStubs(repo)
			m, err := u.GetMovieByID(movie.MovieID)
			tc.checkResponse(movie, m, err)
		})
	}
}

func TestCreateMovie(t *testing.T) {
	now := time.Now()
	arg := model.MovieReq{
		MovieID: "21072001",
		Isbn:    "12131415",
		Title:   "The Lmao Movie",
		Director: model.Director{
			Firstname: "Troai",
			Lastname:  "Lee",
		},
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	testCases := []struct {
		name       string
		buildStubs func(mock_movie *mock_model.MockMovieRepository)
	}{
		{
			name: "OK",
			buildStubs: func(mock_movie *mock_model.MockMovieRepository) {
				mock_movie.EXPECT().Create(gomock.Eq(&arg)).Times(1).Return(nil)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_model.NewMockMovieRepository(ctrl)
			u := NewMovieUseCase(repo)
			tc.buildStubs(repo)
			err := u.CreateMovie(&arg)
			require.NoError(t, err)
		})
	}
}
