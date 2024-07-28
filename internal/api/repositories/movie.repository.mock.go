package repositories

import (
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/mock"
)

type MockMovieRepository struct {
	mock.Mock
	MockBaseRepository[models.MovieModel]
}

func (m *MockMovieRepository) GetPaginated(dto *dtos.PaginationDto) ([]*models.MovieModel, error) {
	args := m.Called()
	return args.Get(0).([]*models.MovieModel), args.Error(1)
}

func (m *MockMovieRepository) GetById(id string) (*models.MovieModel, error) {
	args := m.Called(id)
	return args.Get(0).(*models.MovieModel), args.Error(1)
}

func (m *MockMovieRepository) Create(data *models.MovieModel) (*models.MovieModel, error) {
	args := m.Called(data)
	return data, args.Error(0)
}

func (m *MockMovieRepository) Delete(id string) ( error) {
	args := m.Called(id)
	return args.Error(0)
}

