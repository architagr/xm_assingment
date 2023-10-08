package customerror

type invalidCredentialsError struct {
}

func (err *invalidCredentialsError) Error() string {
	return "invalid credenials"
}

func InitInvalidCredentialsError() error {
	return &invalidCredentialsError{}
}
