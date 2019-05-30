package webrest

import (
	"net/http"
	"regexp"

	"github.com/alekns/go-inversify"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes .
func RegisterRoutes(container inversify.Container, router *gin.RouterGroup) {
	var uuidRegexp = regexp.MustCompile(`^[a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12}$|^[a-f\d]{32}$`)

	// Validate params
	router.Use(func(c *gin.Context) {
		if len(c.Param(userParamID)) > 0 && !uuidRegexp.MatchString(c.Param(userParamID)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "userId has invalid uuid format",
			})
		} else {
			c.Next()
		}
	})

	// Authentication
	RegisterUserControllerRoutes(container, router.Group("/auth/user"))
	RegisterJwtControllerRoutes(container, router.Group("/auth/jwt"))
}
