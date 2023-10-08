package models

import "company_service/enums"

type CreateCompanyRequest struct {
	Name             string                `json:"name" binding:"required"`
	Description      string                `json:"description" binding:"max=3000"`
	NumberOfEmployee int                   `json:"numberOfEmployee" binding:"required,gte=1"`
	Registered       bool                  `json:"registered" binding:"required"`
	Type             enums.CompanyTypeEnum `json:"type" binding:"required,gte=0,lte=3"`
}

type CreateCompanyResponse CompanyDto

func InitCreateCompanyResponse(data *CompanyDto) *CreateCompanyResponse {
	return &CreateCompanyResponse{
		Id:               data.Id,
		Name:             data.Name,
		Description:      data.Description,
		NumberOfEmployee: data.NumberOfEmployee,
		Registered:       data.Registered,
		Type:             data.Type,
	}
}
func (model *CreateCompanyRequest) GetCompanyDTO() *CompanyDto {
	return &CompanyDto{
		Name:             model.Name,
		Description:      model.Description,
		NumberOfEmployee: model.NumberOfEmployee,
		Registered:       model.Registered,
		Type:             model.Type,
	}
}
