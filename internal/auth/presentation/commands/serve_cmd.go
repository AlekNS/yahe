package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alekns/go-inversify"
	authModuleDeps "github.com/alekns/yahe/internal/auth/app/authmodule"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/alekns/yahe/internal/auth/infrastruct"
	"github.com/alekns/yahe/internal/auth/presentation/webrest"
	"github.com/alekns/yahe/internal/helpers"
	"github.com/alekns/yahe/internal/logger"
	"github.com/alekns/yahe/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		var settings = config.GetSettings(helpers.GetSettings("yaheauth", cmd, viper.GetViper()))

		//
		// Prepare logger
		//
		logger.InitLogger(settings.Logger)
		var log = logger.Get("ServeCommand")

		//
		// Process signals
		//
		rootContext, cancelByContext := context.WithCancel(context.Background())
		defer cancelByContext()

		appSignals := make(chan os.Signal, 0)
		signal.Notify(appSignals, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			sig := <-appSignals
			log.Info("stop process. Catch signal", sig)
			cancelByContext()
		}()

		//
		// Top container
		//
		usersStorage, err := infrastruct.NewDBSqlx(settings.Users.StorageURL)
		utils.DieIfError(err, "unable to initialize user storage")
		jwtStorage, err := infrastruct.NewRedis(settings.Jwt.StorageURL)
		utils.DieIfError(err, "unable to initialize jwt storage")

		log.Debug("prepare containers")
		container := inversify.NewContainer("top")
		container.Bind((*config.Settings)(nil)).To(settings)
		container.Bind((*logrus.Logger)(nil)).To(logger.GetRootLogger())

		container.Bind((*sqlx.DB)(nil), "user").To(usersStorage)
		container.Bind((*redis.Client)(nil), "jwt").To(jwtStorage)

		authContainer := inversify.NewContainer("auth")
		authContainer.SetParent(container)

		authModule := authModuleDeps.GetModule()
		authContainer.Load(authModule)

		//
		// Run http server
		//
		if settings.Logger.ConsoleLevel != "debug" {
			gin.SetMode(gin.ReleaseMode)
		} else {
			gin.SetMode(gin.DebugMode)
		}

		var server = webrest.NewHTTPServerImpl(authContainer)

		log.Info("start serve")
		server.Serve(rootContext)
	},
}

// Command .
func Command() *cobra.Command {
	return serveCmd
}
