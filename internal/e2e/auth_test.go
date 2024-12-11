package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/joaops3/go-olist-challenge/internal/api/server"
	"github.com/stretchr/testify/assert"
)


func TestE2EAuth(t *testing.T) {

	server := server.NewHttpServer()
	_, err := server.Serve()
	if err != nil {
		t.Fatalf("could not start server: %v", err)
	}

	t.Run("Teste2eAuth", func(t *testing.T) {
		
	
	
		c := http.Client{}

		r, _ := c.Get("http://localhost:3333/auth")
	
		assert.Equal(t, http.StatusOK, r.StatusCode)
	
		b, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "OK", string(b))
	})

	}