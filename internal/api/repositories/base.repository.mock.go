package repositories

import (
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/mock"
)

type MockBaseRepository[T models.Identifiable] struct {
	mock.Mock
}


func (m *MockBaseRepository[T]) BaseGetPaginated(dto *dtos.PaginationDto) ([]*T, error) {
	args := m.Called()
	return args.Get(0).([]*T), args.Error(1)
}

func (m *MockBaseRepository[T]) BaseGetById(id string) (*T, error) {
	args := m.Called(id)
	return nil, args.Error(0)
}

func (m *MockBaseRepository[T]) BaseCreate(data *T) (*T, error) {
	args := m.Called(data)
	return data, args.Error(0)
}

func (m *MockBaseRepository[T]) BaseDelete(id string) ( error) {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBaseRepository[T]) BaseUpdate(id string, dto any) (bool, error) {
	args := m.Called(id)
	return true, args.Error(0)
}

func (m *MockBaseRepository[T]) BaseSave(data *T) (*T, error) {
	args := m.Called(data)

	if args.Get(0) != nil {
       return args.Get(0).(*T), nil
    }
	return nil, args.Error(1)
}