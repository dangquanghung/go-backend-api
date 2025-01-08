package service

import (
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/repo"
	"github.com/dangquanghung/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct { // private function
	userRepo repo.IUserRepository
	// ...
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. Validate email
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSuccess
}
