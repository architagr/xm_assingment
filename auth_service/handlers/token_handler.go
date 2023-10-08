package handlers

import (
	"auth_service/models"
	"auth_service/services"
)

type ITokenHandler interface {
	GetToken(username, password string) (*models.TokenResponse, error)
}

type tokenHandler struct {
	tokenSvc services.ITokenService
}

func (handler *tokenHandler) GetToken(username, password string) (*models.TokenResponse, error) {
	token, err := handler.tokenSvc.GetToken(username, password)
	if err != nil {
		return nil, err
	}
	return &models.TokenResponse{
		Token: token,
	}, nil
}

func InitTokenHandler(tokenSvc services.ITokenService) ITokenHandler {
	return &tokenHandler{
		tokenSvc: tokenSvc,
	}
}
