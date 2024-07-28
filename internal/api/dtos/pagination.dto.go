package dtos

type PaginationDto struct {
	PageSize int64 `form:"pageSize" json:"pageSize" validate:"required,min=1,max=1000"`
	Current  int64 `form:"current" json:"current" validate:"required,min=1`
}