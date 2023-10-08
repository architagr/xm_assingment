package enums

type CompanyTypeEnum int

const (
	COMPANY_TYPE_CORPORATIONS        CompanyTypeEnum = iota
	COMPANY_TYPE_NON_PROFIT          CompanyTypeEnum = iota
	COMPANY_TYPE_COOPERATIVE         CompanyTypeEnum = iota
	COMPANY_TYPE_SOLE_PROPRIETORSHIP CompanyTypeEnum = iota
)
