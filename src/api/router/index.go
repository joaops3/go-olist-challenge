package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	InitializeRoutes(r)
	PORT := os.Getenv("PORT")
	r.Run(PORT)
}


