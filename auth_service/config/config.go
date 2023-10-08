package config

import (
	customerror "auth_service/custom_error"
	"encoding/json"
	"os"
)

type IConfig interface {
	GetDbConnectionString() string
	GetPortNumber() int
	GetTokenExpiration() int
	GetDatabaseName() string
}

type config struct {
	DbConnectionString string `json:"dbConnectionString" validate:"required,min=3"`
	Port               int    `json:"port" validate:"required,numeric"`
	TokenExpiration    int    `json:"tokenExpiration" validate:"required,numeric"`
	DatabaseName       string `json:"databaseName" validate:"required,min=3"`
}

func (con *config) GetDbConnectionString() string {
	return con.DbConnectionString
}
func (con *config) GetPortNumber() int {
	return con.Port
}
func (con *config) GetTokenExpiration() int {
	return con.TokenExpiration
}
func (con *config) GetDatabaseName() string {
	return con.DatabaseName
}
func InitConfigFromFile(filePath string) (IConfig, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, customerror.InitConfigFileNotFoundError(filePath)
	}
	con := new(config)
	err = json.Unmarshal(fileData, con)
	if err != nil {
		return nil, customerror.InitConfigFileUnmarshalError(err)
	}
	return con, nil
}
