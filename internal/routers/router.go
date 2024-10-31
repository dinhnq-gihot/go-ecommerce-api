package routers

import (
	"go-ecommerce-backend-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		// v1.GET("/ping/:name", Pong)
		v1.GET("/user/1", controllers.NewUserController().GetUserById)
	}
	return r
}
