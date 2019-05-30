package webrest

import (
	"net/http"

	"github.com/alekns/go-inversify"
	authApp "github.com/alekns/yahe/internal/auth/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	userControllerImpl struct {
		logger  *logrus.Entry
		userApp authApp.UserApp
		jwtApp  authApp.JwtApp
	}

	userCtrlCreateRequest struct {
		Name     string `json:"name"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	userCtrlChangePasswordRequest struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
)

func (uc *userControllerImpl) Create(c *gin.Context) {
	var logger = uc.logger.WithField("method", "Create")
	var request = &userCtrlCreateRequest{}

	err := c.BindJSON(request)
	if err != nil {
		ginErrorInvalidRequest(c, logger, err)
		return
	}

	var user = &authApp.User{
		Login:    request.Login,
		Name:     request.Name,
		Password: request.Password,
	}

	user, err = uc.userApp.Create(c.GetHeader(headerTenantID), user)
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("Create was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"login": user.Login,
	})
}

func (uc *userControllerImpl) ActionChangePassword(c *gin.Context) {
	var logger = uc.logger.WithField("method", "ActionChangePassword")
	var request = &userCtrlChangePasswordRequest{}

	err := c.BindJSON(request)
	if err != nil {
		ginErrorInvalidRequest(c, logger, err)
		return
	}

	if request.CurrentPassword == request.NewPassword {
		ginErrorMsg(c, http.StatusUnprocessableEntity, "passwords are same")
		return
	}

	user, err := uc.userApp.ChangePassword(c.GetHeader(headerTenantID),
		c.Param(userParamID),
		request.CurrentPassword, request.NewPassword)
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("Change password was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"login": user.Login,
	})
}

func (uc *userControllerImpl) ActionResetPassword(c *gin.Context) {
	ginErrorMsg(c, http.StatusInternalServerError, "not implemented yet")
}

func (uc *userControllerImpl) ActionInvalidateTokens(c *gin.Context) {
	var logger = uc.logger.WithField("method", "ActionInvalidateTokens")

	err := uc.jwtApp.InvalidateAllTokens(c.GetHeader(headerTenantID), c.Param(userParamID))
	if err != nil {
		if ginAuthErrors(c, err, http.StatusInternalServerError, "") == http.StatusInternalServerError {
			logger.Error("InvalidateAllTokens was failed with:", err)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
	})
}

// RegisterUserControllerRoutes .
func RegisterUserControllerRoutes(container inversify.Container, router *gin.RouterGroup) {
	container.Bind(UserControllerSymbol).ToFactory(func(logger, jwtApp, userApp inversify.Any) (inversify.Any, error) {
		return &userControllerImpl{
			logger:  logger.(*logrus.Logger).WithField("tag", "UserController"),
			jwtApp:  jwtApp.(authApp.JwtApp),
			userApp: userApp.(authApp.UserApp),
		}, nil
	}, (*logrus.Logger)(nil), authApp.JwtAppSymbol, authApp.UserAppSymbol)

	router.POST("", func(c *gin.Context) {
		container.MustGet(UserControllerSymbol).(UserController).Create(c)
	})

	router.GET("/:userId/jwt", func(c *gin.Context) {
		container.MustGet(JwtControllerSymbol).(JwtController).GetIndex(c)
	})

	router.POST("/:userId/actions/change-password", func(c *gin.Context) {
		container.MustGet(UserControllerSymbol).(UserController).ActionChangePassword(c)
	})

	router.POST("/:userId/actions/reset-password", func(c *gin.Context) {
		container.MustGet(UserControllerSymbol).(UserController).ActionResetPassword(c)
	})

	router.POST("/:userId/actions/invalidate-tokens", func(c *gin.Context) {
		container.MustGet(UserControllerSymbol).(UserController).ActionInvalidateTokens(c)
	})
}
