package webrest

import (
	"net/http"

	"github.com/alekns/go-inversify"
	authApp "github.com/alekns/yahe/internal/auth/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	jwtControllerImpl struct {
		logger *logrus.Entry
		jwtApp authApp.JwtApp
	}

	jwtCtrlCreateRequest struct {
		Login    string            `json:"login"`
		Password string            `json:"password"`
		Scope    string            `json:"scope"`
		Attrs    map[string]string `json:"attrs"`
	}

	jwtCtrlValidateRequest struct {
		AccessToken string `json:"accessToken"`
	}
)

func (jc *jwtControllerImpl) GetIndex(c *gin.Context) {
	var logger = jc.logger.WithField("method", "GetIndex")

	result, err := jc.jwtApp.GetAllByUserID(c.GetHeader(headerTenantID), c.Param(userParamID))
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("GetAllByUserID was failed with", err)
		}
		return
	}

	if result == nil {
		ginError(c, http.StatusNotFound)
		return
	}

	response := make([]gin.H, 0, len(result))
	for _, token := range result {
		response = append(response, gin.H{
			"accessToken": token,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (jc *jwtControllerImpl) Create(c *gin.Context) {
	var logger = jc.logger.WithField("method", "Create")

	var request = &jwtCtrlCreateRequest{}

	err := c.BindJSON(request)
	if err != nil {
		ginErrorInvalidRequest(c, logger, err)
		return
	}

	tokenEncoded, err := jc.jwtApp.Create(c.GetHeader(headerTenantID),
		request.Login,
		request.Password,
		&authApp.JwtAppCreateRequest{
			Scope: request.Scope,
			Attrs: request.Attrs,
		})
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("Create was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": tokenEncoded,
	})
}

func (jc *jwtControllerImpl) ActionValidate(c *gin.Context) {
	var logger = jc.logger.WithField("method", "ActionValidate")

	request := &jwtCtrlValidateRequest{}
	err := c.BindJSON(request)
	if err != nil {
		ginErrorInvalidRequest(c, logger, err)
		return
	}

	token, err := jc.jwtApp.Validate(c.GetHeader(headerTenantID), request.AccessToken)
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("Validate was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, token)
}

func (jc *jwtControllerImpl) ActionInvalidate(c *gin.Context) {
	var logger = jc.logger.WithField("method", "ActionInvalidate")

	request := &jwtCtrlValidateRequest{}
	err := c.BindJSON(request)
	if err != nil {
		ginErrorInvalidRequest(c, logger, err)
		return
	}

	err = jc.jwtApp.Invalidate(c.GetHeader(headerTenantID), request.AccessToken)
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("Invalidate was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
	})
}

// RegisterJwtControllerRoutes .
func RegisterJwtControllerRoutes(container inversify.Container, router *gin.RouterGroup) {
	container.Bind(JwtControllerSymbol).ToFactory(func(logger, jwtApp inversify.Any) (inversify.Any, error) {
		return &jwtControllerImpl{
			logger: logger.(*logrus.Logger).WithField("tag", "JwtController"),
			jwtApp: jwtApp.(authApp.JwtApp),
		}, nil
	}, (*logrus.Logger)(nil), authApp.JwtAppSymbol)

	router.POST("", func(c *gin.Context) {
		container.MustGet(JwtControllerSymbol).(JwtController).Create(c)
	})

	router.POST("/actions/validate", func(c *gin.Context) {
		container.MustGet(JwtControllerSymbol).(JwtController).ActionValidate(c)
	})

	router.POST("/actions/invalidate", func(c *gin.Context) {
		container.MustGet(JwtControllerSymbol).(JwtController).ActionInvalidate(c)
	})
}
