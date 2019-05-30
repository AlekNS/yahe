package webrest

import (
	"net/http"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ginErrorMsg(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    msg,
	})
}

func ginError(c *gin.Context, statusCode int) {
	ginErrorMsg(c, statusCode, http.StatusText(statusCode))
}

func ginErrorInvalidRequest(c *gin.Context, logger *logrus.Entry, err error) {
	logger.Error("BindJSON was failed with:", err)
	ginErrorMsg(c, http.StatusBadRequest, "invalid request format")
}

func ginAuthErrors(c *gin.Context, err error, defaultStatusCode int, defaultMsg string) int {
	switch err {
	case app.ErrorNotFound:
		fallthrough
	case app.ErrorUserIsNotActive:
		fallthrough
	case app.ErrorUserAlreadyExists:
		fallthrough
	case app.ErrorUserIsActive:
		fallthrough
	case app.ErrorPasswordIsVeryBasic:
		ginErrorMsg(c, http.StatusUnprocessableEntity, err.Error())
		return http.StatusUnprocessableEntity
	case app.ErrorInvalidJwtToken:
		ginErrorMsg(c, http.StatusBadRequest, err.Error())
		return http.StatusBadRequest
	case app.ErrorPasswordMismatch:
		ginErrorMsg(c, http.StatusBadRequest, "user or password mismatched")
		return http.StatusBadRequest
	case app.ErrorInternalStorageInconsistent:
		ginErrorMsg(c, http.StatusInternalServerError, "internal storage inconsistency")
		return http.StatusInternalServerError
	default:
		if len(defaultMsg) > 0 {
			ginErrorMsg(c, defaultStatusCode, defaultMsg)
		} else {
			ginError(c, defaultStatusCode)
		}
		return defaultStatusCode
	}
}
