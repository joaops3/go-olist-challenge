package repositories

type BaseRepositoryInterface[T any] interface {
	GetPaginated() ([]*T, error)
	GetById(id string) (*T, error)
	Save(v *T) (*T, error)
	Create(v *T) (*T, error)
	Delete(id string) error
}
