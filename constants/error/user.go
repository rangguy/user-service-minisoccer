package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorret     = errors.New("password incorret")
	ErrUsernameExist        = errors.New("username exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorret,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
