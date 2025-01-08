package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// public routers
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.GET("/register")

	}
	// private routers

	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.POST("/active_user")

	}

}
