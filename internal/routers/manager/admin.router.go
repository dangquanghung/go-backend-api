package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	// public routers
	adminRouterPublic := Router.Group("/admin/user")
	{
		adminRouterPublic.GET("/login")

	}
	// private routers

	adminRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())

	{
		adminRouterPrivate.POST("/active_user")

	}

}
