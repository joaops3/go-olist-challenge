package services

import (
	"fmt"

	"github.com/gocarina/gocsv"
	"github.com/joaops3/go-olist-challenge/src/api/dtos"
	"github.com/joaops3/go-olist-challenge/src/api/repositories"
	"github.com/joaops3/go-olist-challenge/src/data/models"
)
type MovieService struct {
	BaseServiceInterface[dtos.CreateMovieDto, dtos.UpdateMovieDto, models.MovieModel]
	Repository repositories.MovieRepositoryInterface
}

func NewMovieService() *MovieService{
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
    
	result, err := s.Repository.Save(movie)
	fmt.Printf("%v",result)
	if err != nil {
		return false, err
	}
	
	return true, nil
}

func (s *MovieService) GetPaginated() ([]*models.MovieModel, error) {

	movies, err := s.Repository.GetPaginated()

	if err != nil {
		return nil, err
	}
	return movies, nil
}