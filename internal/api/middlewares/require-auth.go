package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joaops3/go-olist-challenge/internal/api/repositories"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(403)
		return 
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatus(403)
		return 
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println(err) 
		c.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > claims["exp"].(int64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		id := claims["sub"]

		repo := repositories.NewUserRepository(models.GetDbUserCollection())

		user, err := repo.BaseGetById(id.(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return 
		}

		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 
	}

	c.Next()
}