package internal

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrEmailExists          = errors.New("email already exists")
	ErrValidationFailed     = errors.New("validation failed")
	ErrHashFailed           = errors.New("password hashing failed")
	ErrHashValidationFailed = errors.New("hash validation failed")
	ErrPhoneExists          = errors.New("phone number already exists")
	ErrUserNotActive        = errors.New("user not active")
	ErrFetchingUsers        = errors.New("failed to fetch users")
)
