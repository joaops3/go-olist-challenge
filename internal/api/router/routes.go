package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/internal/api/controllers"
	"github.com/joaops3/go-olist-challenge/internal/api/middlewares"
	"github.com/joaops3/go-olist-challenge/internal/api/services"
)

func InitializeRoutes(r *gin.Engine) {
	
	InitializeMoviesRoutes(r)
	InitializeAuthRoutes(r)
}

func InitializeAuthRoutes(r *gin.Engine){
	service := services.NewAuthService()
	controller := controllers.InitAuthController(service)

	routerGroup := r.Group("/auth")
	routerGroup.POST("/signin", controller.SignIn)
	routerGroup.POST("/signup",controller.SignUp)
}


func InitializeMoviesRoutes(r *gin.Engine){
	service := services.NewMovieService()
	controller := controllers.InitMovieController(service)

	routerGroup := r.Group("/movies")
	routerGroup.POST("/", middlewares.RequireAuth, controller.Post)
	routerGroup.POST("/upload", middlewares.RequireAuth, controller.UploadMovieCsv)
	routerGroup.GET("/", middlewares.RequireAuth, controller.GetPaginated)
	routerGroup.GET("/:id", middlewares.RequireAuth, controller.GetOne)
	routerGroup.PATCH("/:id", middlewares.RequireAuth, controller.Update)
	routerGroup.DELETE("/:id", middlewares.RequireAuth, controller.Delete)
}
