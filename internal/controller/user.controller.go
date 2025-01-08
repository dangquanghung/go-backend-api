package controller

import (
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/service"
	"github.com/dangquanghung/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

// // controller -> service -> repo -> models -> dbs
func (uc *UserController) Register(c *gin.Context) {

	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}
