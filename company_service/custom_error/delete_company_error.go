package customerror

import (
	"fmt"
)

type deleteCompanyError struct {
	companyId string
}

func (err *deleteCompanyError) Error() string {
	return fmt.Sprintf("Error while updating compnay with id %s", err.companyId)
}

func InitDeleteCompanyError(companyId string) error {
	return &deleteCompanyError{
		companyId: companyId,
	}
}
