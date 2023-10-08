package customerror

import (
	"fmt"
)

type companyExistsError struct {
	companyName string
}

func (err *companyExistsError) Error() string {
	return fmt.Sprintf("company with name '%s'already exists", err.companyName)
}

func InitCompanyExistsError(companyName string) error {
	return &companyExistsError{
		companyName: companyName,
	}
}
