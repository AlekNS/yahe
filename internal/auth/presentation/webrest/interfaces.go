package webrest

import (
	"context"

	"github.com/gin-gonic/gin"
)

type (
	// HTTPServer .
	HTTPServer interface {
		// Serve .
		Serve(rootContext context.Context)
	}

	//
	// Controllers
	//

	// JwtController .
	JwtController interface {
		GetIndex(c *gin.Context)
		Create(c *gin.Context)

		ActionValidate(c *gin.Context)
		ActionInvalidate(c *gin.Context)
	}

	// UserController .
	UserController interface {
		Create(c *gin.Context)

		ActionChangePassword(c *gin.Context)
		ActionResetPassword(c *gin.Context)
		ActionInvalidateTokens(c *gin.Context)
	}
)
