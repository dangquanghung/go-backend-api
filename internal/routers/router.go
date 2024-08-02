package routers

import (
	c "github.com/dangquanghung/go-ecommerce-backend-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewUserController().Pong)
		// v1.POST("/ping", pong)
		// v1.PATCH("/ping", pong)
		// v1.PUT("/ping", pong)
		// v1.DELETE("/ping", pong)
	}

	return r
}
