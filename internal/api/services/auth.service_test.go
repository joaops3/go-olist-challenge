package services

import (
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
    suite.Suite
    service AuthServiceInterface
}

func (suite *AuthTestSuite) SetupTest() {
  userRepoMock := new(repositories.MockUserRepository)
  userRepoMock.MockBaseRepository = *new(repositories.MockBaseRepository[models.UserModel])
  suite.service = &AuthService{
	Repository: userRepoMock,
  }
}

func (suite *AuthTestSuite) TestShouldSignUp() {
    assert.Equal(suite.T(), 5, 5)
}




func TestSuite(t *testing.T) {
    suite.Run(t, new(AuthTestSuite))
}