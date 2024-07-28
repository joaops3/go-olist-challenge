package services

import (
	"errors"
	"fmt"

	"github.com/gocarina/gocsv"
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/joaops3/go-olist-challenge/internal/helpers"
)

type MovieService struct {
	Repository repositories.MovieRepositoryInterface
}

type MovieServiceInterface interface {
	BaseServiceInterface[dtos.CreateMovieDto, dtos.UpdateMovieDto, models.MovieModel]
	UploadCsvChunks(file []byte) (bool, error)
	UploadCsv(file []string) (bool, error)
}

func NewMovieService() MovieServiceInterface {
	collection := models.GetDbMovieCollection()
	return &MovieService{Repository: repositories.NewMovieRepository(collection)}
}

func (s *MovieService) UploadCsvChunks(file []byte) (bool, error) {
	var movies []*models.MovieModel
	err := gocsv.UnmarshalBytes(file, &movies)
	if err != nil {
		return false, err
	}
	for _, movie := range movies {
		fmt.Printf("Movie: %+v\n", movie)
	}
	return true, nil
}

func (s *MovieService) UploadCsv(file []string) (bool, error) {
	movie := models.NewMovieModel(file[1], file[2])
	_, err := s.Repository.Create(movie)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *MovieService) GetPaginated(query *dtos.PaginationDto) ([]*models.MovieModel, error) {
	movies, err := s.Repository.GetPaginated(query)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) GetOne(id string) (*models.MovieModel, error) {

	data, err := s.Repository.BaseGetById(id)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *MovieService) Post(dto *dtos.CreateMovieDto) (*models.MovieModel, error) {	

	
	newModel := models.NewMovieModel(dto.Name, dto.Genre)
	
	data, err := s.Repository.BaseSave(newModel)
	
	if err != nil {
		return nil, err
	}
	
	return data, nil
}

func (s *MovieService) Update(id string, dto *dtos.UpdateMovieDto) (*models.MovieModel, error) {

	old, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if old == nil {
		return nil, errors.New("document not found")
	}
	helpers.ObjectAssign(old, dto)
	s.Repository.BaseSave(old)
	
	return old, nil
}

func (s *MovieService) Delete(v string) (bool, error) {

	old, err := s.Repository.GetById(v)
	if err != nil {
		return false, err
	}

	if old == nil {
		return false, errors.New("document not found")
	}

	err = s.Repository.Delete(v)
	if err != nil {
		return false, err
	}
	return true, nil
}