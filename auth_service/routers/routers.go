package routers

import (
	"auth_service/config"
	"auth_service/handlers"
	middlewarePkg "auth_service/middleware"
	"auth_service/models"

	"fmt"
	"net/http"
	xmLogger "xm_logger/logger"

	"github.com/gin-gonic/gin"
)

var ginMiddleware middlewarePkg.IMiddleware[gin.HandlerFunc]

type ginRouter struct {
	ginEngine *gin.Engine
	env       config.IConfig
	logObj    xmLogger.IXmLogger
}

func (router *ginRouter) StartApp(port int) {

	router.ginEngine.Run(fmt.Sprintf(":%d", port))

}

func InitGinRouters(tokenHandler handlers.ITokenHandler, logObj xmLogger.IXmLogger, authConfig config.IConfig) IRouter {
	if tokenHandler == nil {
		panic("handlers not intilized ")
	}
	ginMiddleware = getMiddlewares()
	ginEngine := gin.Default()
	registerInitialCommonMiddleware(ginEngine, logObj)
	routerGroup := getInitialRouteGroup(ginEngine)

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Auth api is working",
		})
	})
	createAuthRoutes(routerGroup, tokenHandler, logObj)
	return &ginRouter{
		ginEngine: ginEngine,
		env:       authConfig,
		logObj:    logObj,
	}
}

func createAuthRoutes(group *gin.RouterGroup, tokenHandler handlers.ITokenHandler, logObj xmLogger.IXmLogger) {

	group.POST("/getToken", func(ginContext *gin.Context) {
		var authRequest models.GetTokenRequest
		if err := ginContext.ShouldBind(&authRequest); err != nil {
			logObj.Error("wrong request body", "error", err)
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypeBind,
			})
			return
		}
		token, err := tokenHandler.GetToken(authRequest.Username, authRequest.Password)
		if err != nil {
			ginContext.Errors = append(ginContext.Errors, &gin.Error{
				Err:  err,
				Type: gin.ErrorTypePrivate,
			})
			return
		}
		ginContext.JSON(http.StatusOK, token)
	})
}
func getInitialRouteGroup(ginEngine *gin.Engine) *gin.RouterGroup {
	return ginEngine.Group("/auth")
}
func registerInitialCommonMiddleware(ginEngine *gin.Engine, logObj xmLogger.IXmLogger) {
	ginEngine.Use(ginMiddleware.GetErrorHandler(logObj))
	ginEngine.Use(ginMiddleware.GetCorsMiddelware())
}
func getMiddlewares() middlewarePkg.IMiddleware[gin.HandlerFunc] {
	return middlewarePkg.InitGinMiddelware()
}
