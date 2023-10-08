package customerror

import (
	"fmt"
)

type updateCompanyError struct {
	companyId string
}

func (err *updateCompanyError) Error() string {
	return fmt.Sprintf("Error while updating compnay with id %s", err.companyId)
}

func InitUpdateCompanyError(companyId string) error {
	return &updateCompanyError{
		companyId: companyId,
	}
}
