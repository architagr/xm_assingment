package customerror

import "fmt"

type configFileUnmarshalError struct {
	err error
}

func (err *configFileUnmarshalError) Error() string {
	return fmt.Sprintf("File Unmarshalling error (%s)", err.err.Error())
}
func InitConfigFileUnmarshalError(err error) error {
	return &configFileUnmarshalError{
		err: err,
	}
}
