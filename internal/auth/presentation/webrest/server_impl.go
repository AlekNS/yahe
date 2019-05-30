package webrest

import (
	"context"
	"net/http"
	"time"

	"github.com/alekns/go-inversify"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type httpServerImpl struct {
	container inversify.Container
	router    *gin.Engine
}

// Serve .
func (hs *httpServerImpl) Serve(rootContext context.Context) {
	settings := hs.container.MustGet((*config.Settings)(nil)).(*config.Settings)

	AuthAPIRouterInit(hs.container, hs.router.Group("/api"))

	var logger = hs.container.MustGet((*logrus.Logger)(nil)).(*logrus.Logger).WithField("tag", "HttpServerImpl")

	var srv = &http.Server{
		Addr:    settings.HTTP.Bind,
		Handler: hs.router,

		// @TODO: From config
		ReadTimeout:       time.Second * 2,
		ReadHeaderTimeout: time.Second * 2,
		WriteTimeout:      time.Second * 2,
		IdleTimeout:       time.Second * 5,
		MaxHeaderBytes:    1 << 15,
	}

	hs.container.Build()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:", err)
		}
	}()

	<-rootContext.Done()

	logger.Info("graceful shutdown")
	ctx, cancel := context.WithTimeout(rootContext, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server shutdown:", err)
	}
}

// NewHTTPServerImpl .
func NewHTTPServerImpl(container inversify.Container) HTTPServer {

	return &httpServerImpl{
		router:    gin.New(),
		container: container,
	}
}
