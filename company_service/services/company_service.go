package services

import (
	"company_service/config"
	customerror "company_service/custom_error"
	"company_service/models"
	"company_service/persistance"
	xmLogger "xm_logger/logger"
)

type ICompanyService interface {
	AddCompany(request *models.CreateCompanyRequest) (*models.CreateCompanyResponse, error)
	Get(request *models.GetCompanyRequest) (*models.GetCompanyResponse, error)
	UpdateCompany(request *models.PatchCompanyRequest) (*models.PatchCompanyResponse, error)
	DeleteCompany(request *models.DeleteCompanyRequest) error
}

type companyService struct {
	configObj   config.IConfig
	loggerObj   xmLogger.IXmLogger
	companyRepo persistance.ICompanyPersistance
}

func InitCompanyService(loggerObj xmLogger.IXmLogger, companyRepo persistance.ICompanyPersistance, configObj config.IConfig) ICompanyService {
	return &companyService{
		loggerObj:   loggerObj,
		configObj:   configObj,
		companyRepo: companyRepo,
	}
}

func (svc *companyService) AddCompany(request *models.CreateCompanyRequest) (*models.CreateCompanyResponse, error) {
	// validate if company already exists
	existingCompany, err := svc.companyRepo.ListCompanies(&models.CompanyFilterModel{
		CompanyName: request.Name,
		PageSize:    1,
		PageNumber:  0,
	})
	if err != nil {
		svc.loggerObj.Error("Error when validating if company already exists", "companyData", request)
		return nil, customerror.InitGenericError()
	}
	if len(existingCompany) > 0 {
		return nil, customerror.InitCompanyExistsError(request.Name)
	}
	// at this point we are sure that the company does not exist with same name
	data := request.GetCompanyDTO()
	id, err := svc.companyRepo.AddCompany(*data)
	if err != nil {
		return nil, err
	}
	data.Id = id
	return models.InitCreateCompanyResponse(data), nil
}
func (svc *companyService) Get(request *models.GetCompanyRequest) (*models.GetCompanyResponse, error) {
	data, err := svc.companyRepo.GetCompany(request.Id)
	if err != nil {
		return nil, err
	}
	return models.InitGetCompanyResponse(data), nil
}
func (svc *companyService) UpdateCompany(request *models.PatchCompanyRequest) (*models.PatchCompanyResponse, error) {
	updateData := request.GetCompanyDTO()

	err := svc.companyRepo.UpdateCompany(*updateData, request.Id)
	if err != nil {
		return nil, customerror.InitUpdateCompanyError(request.Id)
	}
	data, err := svc.companyRepo.GetCompany(request.Id)
	if err != nil {
		return nil, err
	}
	return models.InitPatchCompanyResponse(data), nil
}
func (svc *companyService) DeleteCompany(request *models.DeleteCompanyRequest) error {
	err := svc.companyRepo.DeleteCompany(request.Id)
	if err != nil {
		return customerror.InitDeleteCompanyError(request.Id)
	}
	return nil
}
