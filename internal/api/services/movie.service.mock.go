package services

import (
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/mock"
)

type MockMovieFactory struct {
	mock.Mock
}

func (m *MockMovieFactory) Create(id int) (*models.MovieModel, error) {
	args := m.Called(id)
	return args.Get(0).(*models.MovieModel), args.Error(1)
}