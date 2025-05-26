package service

import (
	"fmt"
	"time"

	"github.com/dangquanghung/go-ecommerce-backend-api/internal/repo"
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/dangquanghung/go-ecommerce-backend-api/internal/utils/random"
	"github.com/dangquanghung/go-ecommerce-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct { // private function
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
	// ...
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)

	// 5. check OTP is avalable

	// 6. user spam...

	// 1. check email exists in db

	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP ...

	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("Otp is :::%d\n", otp)

	// 3. save OTP in Redis with expiration time

	err := us.userAuthRepo.AddOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. send email OTP

	return response.ErrCodeSuccess
}
