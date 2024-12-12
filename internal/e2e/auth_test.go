package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/assert"
)

const BASE_URL = "http://localhost:3333"


func TestE2EAuth(t *testing.T) {
	// config.LoadEnvFile()
	// server := server.NewHttpServer()
	// _, err := server.Serve()
	// if err != nil {
	// 	t.Fatalf("could not start server: %v", err)
	// }



	t.Run("Should signup", func(t *testing.T) {
		url := fmt.Sprintf(BASE_URL + "/auth/signup")
		dto := dtos.SignInDto{Email: "test@gmail.com", Password: "password"}
		body, err := json.Marshal(dto)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
		r, err := http.Post(url,"application/json", bytes.NewBuffer(body))

		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		
		type BodyResponse struct {
			Data  models.JwtResponse
			message string
		}
		
		responseBody := &BodyResponse{}

		err = json.NewDecoder(r.Body).Decode(responseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}
		defer r.Body.Close()
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.NotEmpty(t, responseBody.Data.Token)
		assert.NotEmpty(t, responseBody.Data.Id)
		assert.Equal(t, "operation from controller: success successful", responseBody.message)
	})


	t.Run("authentication login", func(t *testing.T) {
		url := fmt.Sprintf(BASE_URL + "/auth/signin")
		dto := dtos.SignInDto{Email: "test@gmail.com", Password: "password"}
		body, err := json.Marshal(dto)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
		r, err := http.Post(url,"application/json", bytes.NewBuffer(body))

		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		
		type BodyResponse struct {
			Data  models.JwtResponse
			message string
		}
		
		responseBody := &BodyResponse{}

		err = json.NewDecoder(r.Body).Decode(responseBody)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}
		fmt.Println(responseBody)
		defer r.Body.Close()
		assert.Equal(t, http.StatusOK, r.StatusCode)
		assert.NotEmpty(t, responseBody.Data.Token)
		assert.NotEmpty(t, responseBody.Data.Id)

	})

	t.Run("authentication should return 401", func(t *testing.T) {
		c := http.Client{}

		dto := map[string]any{"Email": "test@gmail.com", "Password": "your_password"}

		body, err := json.Marshal(dto)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
		r, err := c.Post("http://localhost:3333/auth/signin", "application/json", bytes.NewBuffer(body))
		
		if err != nil {
				t.Fatalf("failed to request: %v", err)
		}
		defer r.Body.Close()
	
		assert.Equal(t, http.StatusUnauthorized, r.StatusCode)

	})
	

}