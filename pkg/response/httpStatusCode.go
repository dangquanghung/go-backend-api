package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid

	ErrInvalidToken = 30001 // token is invalid
	ErrInvalidOTP   = 3002
	// Register Code
	ErrCodeUserHasExists = 50001 // user already registered
)

// message

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
	ErrInvalidOTP:       "Otp error",

	ErrCodeUserHasExists: "user has already been registered",
}
