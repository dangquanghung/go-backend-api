package controller

import (
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{ //map string
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"dangquanghung", "cr7", "m10"},
	})
}
