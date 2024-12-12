package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"github.com/stretchr/testify/assert"
)

const BASE_URL = "http://localhost:3333"


func TestE2EAuth(t *testing.T) {


	getToken := func() string {
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
		return responseBody.Data.Token
	}
	


	t.Run("Should upload csv", func(t *testing.T) {
		token := getToken()

		url := fmt.Sprintf(BASE_URL + "/movies/upload")

		var b bytes.Buffer
		w := multipart.NewWriter(&b)
	
		fw, err := w.CreateFormFile("file", "test.csv")
		if err != nil {
			t.Fatalf("failed to create form file: %v", err)
		}

		fileContent := []byte("id,name,genre\n1,Movie1,Action\n2,Movie2,Drama")
		_, err = fw.Write(fileContent)
		if err != nil {
			t.Fatalf("failed to write to form file: %v", err)
		}

		w.Close()
	
		req, err := http.NewRequest("POST", url, &b)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", w.FormDataContentType())

		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+ token)

		client := &http.Client{}
		r, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		assert.Equal(t, http.StatusOK, r.StatusCode)

	})
		



}