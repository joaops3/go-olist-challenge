package repositories

import (
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
	MockBaseRepository[models.UserModel]
}



func (m *MockUserRepository) GetByEmail(id string) (*models.UserModel, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.UserModel), nil
	}
	return nil, args.Error(1)
}



