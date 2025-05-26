package repo

import (
	"fmt"
	"time"

	"github.com/dangquanghung/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email) //usr:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
