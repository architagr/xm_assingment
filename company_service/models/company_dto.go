package models

import (
	"company_service/enums"
)

type CompanyDto struct {
	Id               interface{}           `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string                `json:"name" bson:"name"`
	Description      string                `json:"description,omitempty" bson:"description,omitempty"`
	NumberOfEmployee int                   `json:"numberOfEmployee" bson:"numberOfEmployee"`
	Registered       bool                  `json:"registered" bson:"registered"`
	Type             enums.CompanyTypeEnum `json:"type" bson:"type"`
}
