package dtos

import validation "github.com/go-ozzo/ozzo-validation"

type SignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *SignInDto) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Email, validation.Required.Error("O Email é obrigatório"), validation.Length(1, 50)),
		validation.Field(&d.Password, validation.Required.Error("A Senha é obrigatório"), validation.Length(1, 50)),
	)
}