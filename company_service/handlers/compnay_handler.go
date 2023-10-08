package handlers

import (
	"company_service/models"
	"company_service/services"
	"fmt"
	"net/http"
	xmLogger "xm_logger/logger"

	"github.com/gin-gonic/gin"
)

type ICompanyHandler interface {
	Add(ginContext *gin.Context)
	Get(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
}

func InitCompanyHandler(loggerObj xmLogger.IXmLogger, companyService services.ICompanyService) ICompanyHandler {
	return &companyHandler{
		loggerObj:      loggerObj,
		companyService: companyService,
	}
}

type companyHandler struct {
	loggerObj      xmLogger.IXmLogger
	companyService services.ICompanyService
}

func (ctrl *companyHandler) Add(ginContext *gin.Context) {
	var createRequest models.CreateCompanyRequest
	if err := ginContext.ShouldBind(&createRequest); err != nil {
		ctrl.loggerObj.Error("wrong request body", "error", err)
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
		})
		return
	}
	data, err := ctrl.companyService.AddCompany(&createRequest)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypePrivate,
		})
		return
	}
	ginContext.JSON(http.StatusCreated, data)
}
func (ctrl *companyHandler) Get(ginContext *gin.Context) {
	id, ok := ginContext.Params.Get("id")
	if !ok {
		ctrl.loggerObj.Error("wrong request URI")
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  fmt.Errorf("request param is not present"),
			Type: gin.ErrorTypeBind,
		})
		return
	}
	var getRequest = models.GetCompanyRequest{
		Id: id,
	}
	data, err := ctrl.companyService.Get(&getRequest)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypePrivate,
		})
		return
	}
	ginContext.JSON(http.StatusOK, data)
}
func (ctrl *companyHandler) Update(ginContext *gin.Context) {
	var patchRequest models.PatchCompanyRequest
	if err := ginContext.ShouldBind(&patchRequest); err != nil {
		ctrl.loggerObj.Error("wrong request body", "error", err)
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
		})
		return
	}
	id, ok := ginContext.Params.Get("id")
	if !ok {
		ctrl.loggerObj.Error("wrong request URI")
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  fmt.Errorf("request param is not present"),
			Type: gin.ErrorTypeBind,
		})
		return
	}
	patchRequest.Id = id
	data, err := ctrl.companyService.UpdateCompany(&patchRequest)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypePrivate,
		})
		return
	}
	ginContext.JSON(http.StatusOK, data)
}
func (ctrl *companyHandler) Delete(ginContext *gin.Context) {
	id, ok := ginContext.Params.Get("id")
	if !ok {
		ctrl.loggerObj.Error("wrong request URI")
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  fmt.Errorf("request param is not present"),
			Type: gin.ErrorTypeBind,
		})
		return
	}
	var deleteRequest = models.DeleteCompanyRequest{
		Id: id,
	}

	err := ctrl.companyService.DeleteCompany(&deleteRequest)
	if err != nil {
		ginContext.Errors = append(ginContext.Errors, &gin.Error{
			Err:  err,
			Type: gin.ErrorTypePrivate,
		})
		return
	}
	ginContext.JSON(http.StatusNoContent, nil)
}
