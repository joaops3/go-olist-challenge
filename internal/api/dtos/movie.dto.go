package dtos

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateMovieDto struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}


func (d *CreateMovieDto) Validate() error{
	return validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required.Error("O Nome é obrigatório"), validation.Length(1, 50)),
	)
}

type UpdateMovieDto struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}