package customerror

import "fmt"

type userNameNotFoundError struct {
	userName string
}

func (err *userNameNotFoundError) Error() string {
	return fmt.Sprintf("no user found with username: %s", err.userName)
}

func InitUserNameNotFoundError(username string) error {
	return &userNameNotFoundError{
		userName: username,
	}
}
