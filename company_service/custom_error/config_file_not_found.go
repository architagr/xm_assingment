package customerror

import "fmt"

type configFileNotFound struct {
	filePath string
}

func (err *configFileNotFound) Error() string {
	return fmt.Sprintf("File (%s) is not found", err.filePath)
}
func InitConfigFileNotFoundError(filePath string) error {
	return &configFileNotFound{
		filePath: filePath,
	}
}
