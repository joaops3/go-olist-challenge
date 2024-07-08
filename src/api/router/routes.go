package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/src/api/controllers"
	"github.com/joaops3/go-olist-challenge/src/api/services"
)

func InitializeRoutes(r *gin.Engine) {
	
	InitializeMoviesRoutes(r)
}


func InitializeMoviesRoutes(r *gin.Engine){
	service := services.NewMovieService()
	controller := controllers.InitMovieController(*service)
	routerGroup := r.Group("/movies")
	routerGroup.POST("/",controller.Post)
	routerGroup.POST("/upload",controller.UploadMovieCsv)
	routerGroup.GET("/",controller.GetPaginated)
	routerGroup.GET("/:id",controller.GetOne)
	routerGroup.PATCH("/:id",controller.Update)
	routerGroup.DELETE("/:id",controller.Delete)
	
}

