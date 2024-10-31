package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"id":      id,
		"name":    name,
		"users": []string{
			"dinhnq",
			"dinhnq2",
			"dinhnq3",
		},
	})
}
