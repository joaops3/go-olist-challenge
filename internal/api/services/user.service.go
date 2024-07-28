package services

import (
	"errors"
	"mime/multipart"

	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
)

type userService struct {
	Repository repositories.UserRepositoryInterface
	FileService FileServiceInterface
}

type UserServiceInterface interface {
	UploadPhoto(user *models.UserModel, file *multipart.FileHeader) (*models.UserModel, error)
}

func NewUserService() UserServiceInterface {
	collection := models.GetDbUserCollection()
	fileService := NewFileService()
	return &userService{
		Repository: repositories.NewUserRepository(collection),
		FileService: fileService,
	}
}

func (s *userService)UploadPhoto(user *models.UserModel, file *multipart.FileHeader) (*models.UserModel, error) {
	newFile, err := s.FileService.GenerateFile(file, BUCKET_GENERIC, models.FILE_TYPE_GENERIC)

	if err != nil {
		return nil, err
	}

	if user.ProfileImg != nil {
		err = s.FileService.DeleteFile(string(BUCKET_GENERIC), user.ProfileImg.Key)
		if err != nil {
			return nil, errors.New("error deleting")
		}
	
	}

	user.ProfileImg = newFile

	data, err := s.Repository.BaseSave(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}


