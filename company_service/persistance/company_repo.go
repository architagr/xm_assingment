package persistance

import (
	customerror "company_service/custom_error"
	"company_service/filters"
	"company_service/models"
	xmLogger "xm_logger/logger"

	"go.mongodb.org/mongo-driver/bson"
)

type ICompanyPersistance interface {
	AddCompany(data models.CompanyDto) (interface{}, error)
	DeleteCompany(id string) error
	UpdateCompany(data models.CompanyDto, id string) error
	GetCompany(id string) (*models.CompanyDto, error)
	ListCompanies(filter *models.CompanyFilterModel) ([]models.CompanyDto, error)
}

type companyPersistance struct {
	loggerObj     xmLogger.IXmLogger
	collectionObj ICollection[models.CompanyDto]
}

func InitCompanyPersistance(loggerObj xmLogger.IXmLogger,
	conn IConnection, databaseName string) (ICompanyPersistance, error) {
	collObj, err := InitCollection[models.CompanyDto](conn, databaseName, "companies")
	if err != nil {
		return nil, err
	}
	return &companyPersistance{
		loggerObj:     loggerObj,
		collectionObj: collObj,
	}, nil

}

func (repo *companyPersistance) AddCompany(data models.CompanyDto) (interface{}, error) {
	repo.loggerObj.Debug("adding new company", "companyInfo", data)

	id, err := repo.collectionObj.AddSingle(data)
	if err != nil {
		repo.loggerObj.Error("Error when adding a new company", "dbError", err)
		return nil, customerror.InitAddCompanyError(data.Name, data.Type)
	}
	return id, nil
}

func (repo *companyPersistance) DeleteCompany(id string) error {
	return repo.collectionObj.Delete(id)
}

func (repo *companyPersistance) UpdateCompany(data models.CompanyDto, id string) error {
	return repo.collectionObj.UpdateSingle(data, id)
}

func (repo *companyPersistance) GetCompany(id string) (*models.CompanyDto, error) {
	return repo.collectionObj.GetById(id)
}

func (repo *companyPersistance) ListCompanies(filterData *models.CompanyFilterModel) (list []models.CompanyDto, err error) {
	var filter filters.IFilter = nil
	_filter := bson.M{}
	if filterData != nil {
		if filterData.CompanyName != "" {
			filter = filters.InitCompanynameFilter(filter, filters.AND, filters.EQUAL, filterData.CompanyName)
		}
	}
	if filter != nil {
		_filter = filter.Build()
	}
	return repo.collectionObj.Get(_filter, int64(filterData.PageSize), int64(filterData.PageNumber))
}
