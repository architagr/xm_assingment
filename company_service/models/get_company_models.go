package models

type GetCompanyRequest struct {
	Id string `json:"id" uri:"id"`
}

type GetCompanyResponse CompanyDto

func InitGetCompanyResponse(data *CompanyDto) *GetCompanyResponse {
	return &GetCompanyResponse{
		Id:               data.Id,
		Name:             data.Name,
		Description:      data.Description,
		NumberOfEmployee: data.NumberOfEmployee,
		Registered:       data.Registered,
		Type:             data.Type,
	}
}
