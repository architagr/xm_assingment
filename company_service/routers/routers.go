package routers

import (
	"company_service/config"
	"company_service/handlers"
	middlewarePkg "company_service/middleware"

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

func InitGinRouters(companyHandler handlers.ICompanyHandler, logObj xmLogger.IXmLogger, authConfig config.IConfig) IRouter {
	if companyHandler == nil {
		panic("handlers not intilized ")
	}
	ginMiddleware = getMiddlewares()
	ginEngine := gin.Default()
	registerInitialCommonMiddleware(ginEngine, logObj)
	routerGroup := getInitialRouteGroup(ginEngine)

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Company api is working",
		})
	})
	createAuthRoutes(routerGroup, companyHandler, logObj)
	return &ginRouter{
		ginEngine: ginEngine,
		env:       authConfig,
		logObj:    logObj,
	}
}

func createAuthRoutes(group *gin.RouterGroup, companyHandler handlers.ICompanyHandler, logObj xmLogger.IXmLogger) {
	group.POST("/", ginMiddleware.GetAuthHandler(), companyHandler.Add)
	group.PATCH("/:id", ginMiddleware.GetAuthHandler(), companyHandler.Update)
	group.DELETE("/:id", ginMiddleware.GetAuthHandler(), companyHandler.Delete)
	group.GET("/:id", companyHandler.Get)
}
func getInitialRouteGroup(ginEngine *gin.Engine) *gin.RouterGroup {
	return ginEngine.Group("/company")
}
func registerInitialCommonMiddleware(ginEngine *gin.Engine, logObj xmLogger.IXmLogger) {
	ginEngine.Use(ginMiddleware.GetErrorHandler(logObj))
	ginEngine.Use(ginMiddleware.GetCorsMiddelware())
}
func getMiddlewares() middlewarePkg.IMiddleware[gin.HandlerFunc] {
	return middlewarePkg.InitGinMiddelware()
}
