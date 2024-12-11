package services

import (
	"errors"
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {

	t.Run("Create Movie", func(t *testing.T) {
		
		mockRepo := new(repositories.MockMovieRepository)
		// Initialize the base repository mock
		mockRepo.MockBaseRepository = *new(repositories.MockBaseRepository[models.MovieModel])

		// Mocked input (ensure it's a pointer)
		input := models.NewMovieModel("Test", "Test")
		mockRepo.MockBaseRepository.On("BaseSave", mock.AnythingOfType("*models.MovieModel")).Return(input, nil)

		service := &MovieService{
			Repository: mockRepo,
		}

		dto := dtos.CreateMovieDto{
			Name:  "Test",
			Genre: "Test",
		}

		output, err := service.Post(&dto)

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "Test", output.Name)
		assert.Equal(t, "Test", output.Genre)
		mockRepo.AssertExpectations(t)
	})
	


	t.Run("Create Movie Fail", func(t *testing.T) {

		mockRepo := new(repositories.MockMovieRepository)
		mockRepo.MockBaseRepository = *new(repositories.MockBaseRepository[models.MovieModel])

		
		mockRepo.MockBaseRepository.On("BaseSave", mock.AnythingOfType("*models.MovieModel")).Return(nil, errors.New("error"))

		service := &MovieService{
			Repository: mockRepo,
		}

		dto := dtos.CreateMovieDto{
			Name:  "Test",
			Genre: "Test",
		}

		_, err := service.Post(&dto)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("error"), err)
		mockRepo.AssertExpectations(t)
	})
}


