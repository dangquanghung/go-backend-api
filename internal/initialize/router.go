package initialize

import (
	"fmt"

	c "github.com/dangquanghung/go-ecommerce-backend-api/internal/controller"
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Before -->AA")
		c.Next()
		fmt.Println("After -->AA")
	}

}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before-->BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before-->CC")
	c.Next()
	fmt.Println("After -->CC")
}
func InitRouter() *gin.Engine {
	r := gin.Default()

	// use middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewUserController().GetUserById)
		// v1.POST("/ping", pong)
		// v1.PATCH("/ping", pong)
		// v1.PUT("/ping", pong)
		// v1.DELETE("/ping", pong)
	}

	return r
}
