package webrest

import (
	"net/http"
	"strings"

	"github.com/alekns/go-inversify"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/gin-gonic/gin"
)

func defaultMiddleware(c *gin.Context) {
	if len(c.GetHeader("x-tenant-id")) == 0 || len(c.GetHeader("x-tenant-id")) > 36 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "X-tenant-id header is not valid or not set up",
		})
		return
	}

	if !strings.Contains(c.ContentType(), "application/json") {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Content-Type header is not acceptable",
		})
		return
	}

	if !strings.Contains(c.GetHeader("accept"), "application/json") {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Accept header is not acceptable",
		})
		return
	}

	c.Next()
}

// AuthAPIRouterInit .
func AuthAPIRouterInit(container inversify.Container, router *gin.RouterGroup) error {
	settings := container.MustGet((*config.Settings)(nil)).(*config.Settings)

	var authGroup = router.Group("/v1")

	if settings.HTTP.LogAccess {
		authGroup.Use(gin.Logger())
	}

	authGroup.Use(gin.Recovery())
	authGroup.Use(defaultMiddleware)

	RegisterRoutes(container, authGroup)

	return nil
}
