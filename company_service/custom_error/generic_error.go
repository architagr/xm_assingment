package customerror

type genericError struct {
}

func (err *genericError) Error() string {
	return "something went wrong! please contact admin."
}

func InitGenericError() error {
	return &genericError{}
}
