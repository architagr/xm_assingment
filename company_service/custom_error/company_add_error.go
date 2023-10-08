package customerror

import (
	"company_service/enums"
	"fmt"
)

type addCompanyError struct {
	companyName string
	companyType enums.CompanyTypeEnum
}

func (err *addCompanyError) Error() string {
	return fmt.Sprintf("Error in adding new company (name: %s, of type: %d)", err.companyName, err.companyType)
}

func InitAddCompanyError(companyName string,
	companyType enums.CompanyTypeEnum) error {
	return &addCompanyError{
		companyName: companyName,
		companyType: companyType,
	}
}
