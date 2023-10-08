package customerror

import "fmt"

type wrongSigningMethodError struct {
	alg interface{}
}

func (err *wrongSigningMethodError) Error() string {
	return fmt.Sprintf("unexpected signing method: %v", err.alg)
}

func InitWrongSigningMethodError(alg interface{}) error {
	return &wrongSigningMethodError{
		alg: alg,
	}
}
