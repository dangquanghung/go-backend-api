package controller

import (
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) Pong(c *gin.Context) {

	// response.SuccessResponse(c, 20001, []string{"dangquanghung", "cr7"})
	// response.ErrorResponse(c, 20003, "no need!")
}
