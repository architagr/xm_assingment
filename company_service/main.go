package main

import (
	"company_service/config"
	"company_service/handlers"
	"company_service/persistance"
	"company_service/routers"
	"company_service/services"
	"flag"
	"io"
	"log/slog"
	"os"
	xmJsonLogger "xm_logger/json_logger"
	xmLogger "xm_logger/logger"
)

var xmLoggerObj xmLogger.IXmLogger
var authConfig config.IConfig
var conn persistance.IConnection
var companyRepo persistance.ICompanyPersistance
var companyService services.ICompanyService
var companyHandler handlers.ICompanyHandler
var (
	configFile  = flag.String("configFile", "config.json", "this value is path of the config file")
	useFileLog  = flag.Bool("useFileLog", true, "this is used to have logs in file")
	logFilePath = flag.String("logFilePath", "companyServiceLogs.txt", "this value is path of the logs")
	port        = flag.Int("port", 8081, "this value is used when we run the service on local")
)

func main() {
	flag.Parse()
	InitLogger()
	InitConfig()
	InitConnection()
	InitRepo()
	InitService()
	InitHandlers()
	xmLoggerObj.Debug("service is running", "port", port)
	routers.InitGinRouters(companyHandler, xmLoggerObj, authConfig).StartApp(*port)
}
func InitHandlers() {
	companyHandler = handlers.InitCompanyHandler(xmLoggerObj, companyService)
}
func InitService() {
	companyService = services.InitCompanyService(xmLoggerObj, companyRepo, authConfig)
}
func InitRepo() {
	tokenRepoObj, err := persistance.InitCompanyPersistance(xmLoggerObj, conn, authConfig.GetDatabaseName())
	if err != nil {
		xmLoggerObj.Error("error initilizing Company Repository", "error", err)
		panic(err)
	}
	companyRepo = tokenRepoObj
}
func InitConfig() {
	con, err := config.InitConfigFromFile(*configFile)
	if err != nil {
		xmLoggerObj.Error("error in loading configurations", "error", err)
		panic(err)
	}
	authConfig = con
}
func InitConnection() {
	connObj, err := persistance.InitConnection(authConfig.GetDbConnectionString(), 120)
	if err != nil {
		xmLoggerObj.Error("error in connecting to db", "error", err)
		panic(err)
	}
	conn = connObj
}
func InitLogger() {

	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)
	var w io.Writer
	if *useFileLog {
		w, _ = os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		w = os.Stdout
	}
	xmLoggerObj = xmJsonLogger.InitXmJsonLogger(w, &slog.HandlerOptions{Level: programLevel}, "companyService")
}
