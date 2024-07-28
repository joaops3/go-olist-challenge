package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository repositories.UserRepositoryInterface
}

type AuthServiceInterface interface {
	SignIn(dto *dtos.SignInDto) (*models.JwtResponse, error)
	SignUp(dto *dtos.CreateUserDto) (*models.JwtResponse, error)
}

func NewAuthService() AuthServiceInterface {
	collection := models.GetDbUserCollection()
	return &AuthService{Repository: repositories.NewUserRepository(collection)}
}


func (r *AuthService)SignIn(dto *dtos.SignInDto) (*models.JwtResponse, error) {
	found, err := r.Repository.GetByEmail(dto.Email) 

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(dto.Password))

	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	token, err := generateTokenJWT(found.ID.Hex())
	if err != nil {
		return nil, err
	}

	resp := models.JwtResponse{
		Token: token,
		Id: found.ID.Hex(),
	}
	
	return &resp, nil
}

func (r *AuthService)SignUp(dto *dtos.CreateUserDto) (*models.JwtResponse, error) {

	found, err := r.Repository.GetByEmail(dto.Email) 

	if err != nil {
		return nil, err
	}

	if found != nil {
		return nil, errors.New("email is already being used")
	}
	
	
	model := models.NewUserModel(dto.Email, dto.Password)
	


	hashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return nil, err
	}
	model.Password = string(hashed)

	data, err := r.Repository.BaseCreate(model)
	if err != nil {
		return nil, err
	}

	token, err := generateTokenJWT(data.ID.Hex())
	if err != nil {
		return nil, err
	}

	resp := models.JwtResponse{
		Token: token,
		Id: data.ID.Hex(),
	}

	return &resp, nil
}


func generateTokenJWT(id string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	
	
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, err
}