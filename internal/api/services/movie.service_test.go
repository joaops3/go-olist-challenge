package services

import (
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMovie(t *testing.T) {
	mockRepo := new(repositories.MockMovieRepository)

	mockRepo.On("Create", mock.AnythingOfType("*models.Movie")).Return(nil)


	input := models.NewMovieModel("Test", "Test")

	service := &MovieService{
		Repository: mockRepo,
	}

	output, err := service.Repository.Create(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockRepo.AssertExpectations(t)
}