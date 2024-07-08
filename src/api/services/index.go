package services

type BaseServiceInterface[DtoCreate any, DtoUpdate any, T any] interface {
	GetPaginated() ([]T, error)
	GetOne(id string) T
	Post(v DtoCreate) T
	Update(dtoUpdate DtoUpdate) T
	Delete(v T) T
}