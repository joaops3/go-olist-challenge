package services

import "github.com/joaops3/go-olist-challenge/internal/api/dtos"

type BaseServiceInterface[DtoCreate any, DtoUpdate any, T any] interface {
	GetPaginated(query *dtos.PaginationDto) ([]*T, error)
	GetOne(id string) (*T, error)
	Post(v *DtoCreate) (*T, error)
	Update(id string, dtoUpdate *DtoUpdate) (*T, error)
	Delete(v string) (bool, error)
}