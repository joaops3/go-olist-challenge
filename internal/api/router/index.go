package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() (*gin.Engine, error) {
	r := gin.Default()
	InitializeRoutes(r)
	PORT := os.Getenv("PORT")
	err := r.Run(PORT)
	
	if err != nil {
		return nil, err
	}
	return r, nil
}


