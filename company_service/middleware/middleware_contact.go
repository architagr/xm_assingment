package middleware

import xmLogger "xm_logger/logger"

type IMiddleware[T any] interface {
	GetCorsMiddelware() T
	GetErrorHandler(logObj xmLogger.IXmLogger) T
	GetAuthHandler() T
}
