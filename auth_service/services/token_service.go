package services

import (
	"auth_service/config"
	customerror "auth_service/custom_error"
	"auth_service/persistance"
	"time"
	xmLogger "xm_logger/logger"
	"xm_token/token"

	"github.com/dgrijalva/jwt-go"
)

type ITokenService interface {
	GetToken(username, password string) (string, error)
}

type tokenService struct {
	configObj config.IConfig
	loggerObj xmLogger.IXmLogger
	tokenRepo persistance.ITokenPersistance
}

func InitTokenService(loggerObj xmLogger.IXmLogger,
	tokenRepo persistance.ITokenPersistance, configObj config.IConfig) ITokenService {
	return &tokenService{
		loggerObj: loggerObj,
		configObj: configObj,
		tokenRepo: tokenRepo,
	}
}

func (svc *tokenService) GetToken(username, password string) (string, error) {

	user, err := svc.tokenRepo.GetUser(username)
	if err != nil {
		return "", err
	}

	ok := user.Password.CheckPasswordHash(password)
	if !ok {
		return "", customerror.InitInvalidCredentialsError()
	}
	svc.loggerObj.Debug("expiration time", "time", svc.configObj.GetTokenExpiration())
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(svc.configObj.GetTokenExpiration()) * time.Minute)

	return token.GenrateToken(&token.JwtClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}, expirationTime)
}
