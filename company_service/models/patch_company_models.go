package models

type PatchCompanyRequest struct {
	Id string `json:"id" uri:"id"`
	CreateCompanyRequest
}
type PatchCompanyResponse CompanyDto

func (model *PatchCompanyRequest) GetCompanyDTO() *CompanyDto {
	return &CompanyDto{
		Name:             model.Name,
		Description:      model.Description,
		NumberOfEmployee: model.NumberOfEmployee,
		Registered:       model.Registered,
		Type:             model.Type,
	}
}

func InitPatchCompanyResponse(data *CompanyDto) *PatchCompanyResponse {
	return &PatchCompanyResponse{
		Id:               data.Id,
		Name:             data.Name,
		Description:      data.Description,
		NumberOfEmployee: data.NumberOfEmployee,
		Registered:       data.Registered,
		Type:             data.Type,
	}
}
