package main

import (
	"auth_service/config"
	"auth_service/handlers"
	"auth_service/persistance"
	"auth_service/routers"
	"auth_service/services"
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
var tokenRepo persistance.ITokenPersistance
var tokenService services.ITokenService
var tokenHandler handlers.ITokenHandler
var (
	configFile  = flag.String("configFile", "config.json", "this value is path of the config file")
	useFileLog  = flag.Bool("useFileLog", true, "this is used to have logs in file")
	logFilePath = flag.String("logFilePath", "authServiceLogs.txt", "this value is path of the logs")
	port        = flag.Int("port", 8080, "this value is used when we run the service on local")
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
	routers.InitGinRouters(tokenHandler, xmLoggerObj, authConfig).StartApp(*port)
}
func InitHandlers() {
	tokenHandler = handlers.InitTokenHandler(tokenService)
}
func InitService() {
	tokenService = services.InitTokenService(xmLoggerObj, tokenRepo, authConfig)
}
func InitRepo() {
	tokenRepoObj, err := persistance.InitTokenPersistance(xmLoggerObj, conn, authConfig.GetDatabaseName())
	if err != nil {
		xmLoggerObj.Error("error initilizing Token Repository", "error", err)
		panic(err)
	}
	tokenRepo = tokenRepoObj
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
	xmLoggerObj = xmJsonLogger.InitXmJsonLogger(w, &slog.HandlerOptions{Level: programLevel}, "authService")
}
