package user

import (
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routers
	// this is non-dependency
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependency := controller.NewUserController(us)

	userController, _ := wire.InitUserRouterHandler()

	// WIRE go

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register) // register => yes => no
		userRouterPublic.POST("/otp")                               //

	}

	// private routers

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.GET("/get_info")

	}

}
