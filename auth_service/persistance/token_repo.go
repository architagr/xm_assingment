package persistance

import (
	customerror "auth_service/custom_error"
	"auth_service/models"
	xmLogger "xm_logger/logger"

	"go.mongodb.org/mongo-driver/bson"
)

type ITokenPersistance interface {
	GetUser(username string) (*models.UserModel, error)
}

type tokenPersistance struct {
	loggerObj     xmLogger.IXmLogger
	collectionObj ICollection[models.UserModel]
}

func InitTokenPersistance(loggerObj xmLogger.IXmLogger,
	conn IConnection, databaseName string) (ITokenPersistance, error) {
	collObj, err := InitCollection[models.UserModel](conn, databaseName, "users")
	if err != nil {
		return nil, err
	}
	return &tokenPersistance{
		loggerObj:     loggerObj,
		collectionObj: collObj,
	}, nil

}
func (repo *tokenPersistance) GetUser(username string) (*models.UserModel, error) {
	repo.loggerObj.Debug("getting user with username", "username", username)
	filter := bson.M{"userName": username}

	user, err := repo.collectionObj.Get(filter)
	if err != nil {
		repo.loggerObj.Error("error when getting user with username", "username", username, "dbError", err)
		err = customerror.InitUserNameNotFoundError(username)
	}
	return user, err
}
