package authmodule

import (
	"github.com/alekns/go-inversify"
	auth "github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/internal/auth/app/app"
	"github.com/alekns/yahe/internal/auth/app/services"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

// GetModule .
func GetModule() *inversify.Module {
	return inversify.NewModule("auth").
		Register(func(c inversify.ContainerBinder) error {
			//
			// Services
			//

			c.Bind(auth.UserRepositoryServiceSymbol).ToFactory(func(sqlxDbArg inversify.Any) (inversify.Any, error) {
				sqlxDb := sqlxDbArg.(*sqlx.DB)
				return services.NewUserRepositoryServiceSqlx(sqlxDb), nil
			}, (*sqlx.DB)(nil))

			c.Bind(auth.JwtRepositoryServiceSymbol).ToFactory(func(conf, redisConn inversify.Any) (inversify.Any, error) {
				settings := conf.(*config.Settings)
				redisClient := redisConn.(*redis.Client)
				return services.NewJwtRepositoryServiceRedis(settings.Jwt.KeysPrefix, redisClient), nil
			}, (*config.Settings)(nil), (*redis.Client)(nil))

			c.Bind(auth.PasswordServiceSymbol).ToFactory(func(conf inversify.Any) (inversify.Any, error) {
				// settings := conf.(*config.Settings)
				return services.NewPasswordServiceBcrypt(16), nil
			}, (*config.Settings)(nil))

			c.Bind(auth.JwtServiceSymbol).ToFactory(func(conf inversify.Any) (inversify.Any, error) {
				settings := conf.(*config.Settings)
				return services.NewJwtServiceImpl(settings), nil
			}, (*config.Settings)(nil))

			//
			// Events
			//

			c.Bind(auth.DomainEventsSymbol).ToFactory(func() (inversify.Any, error) {
				return services.NewSyncEventsImpl(), nil
			}).InSingletonScope()

			//
			// Applications
			//

			c.Bind(auth.UserAppSymbol).ToFactory(func(conf, events, userRep, passwdSvc inversify.Any) (inversify.Any, error) {
				settings := conf.(*config.Settings)
				domainEvents := events.(auth.DomainEvents)
				userRepository := userRep.(auth.UserRepositoryService)
				passwordService := passwdSvc.(auth.PasswordService)

				return app.NewUserApp(
					settings,
					domainEvents,
					userRepository,
					passwordService), nil
			}, (*config.Settings)(nil), auth.DomainEventsSymbol, auth.UserRepositoryServiceSymbol, auth.PasswordServiceSymbol)

			c.Bind(auth.JwtAppSymbol).ToFactory(func(conf, userApp, jwtRep, jwtSvc inversify.Any) (inversify.Any, error) {
				settings := conf.(*config.Settings)
				userApplication := userApp.(auth.UserApp)
				jwtRepository := jwtRep.(auth.JwtRepositoryService)
				jwtService := jwtSvc.(auth.JwtService)

				return app.NewJwtApp(settings,
					userApplication,
					jwtRepository,
					jwtService), nil
			}, (*config.Settings)(nil), auth.UserAppSymbol, auth.JwtRepositoryServiceSymbol, auth.JwtServiceSymbol)

			return nil
		}).
		UnRegister(func(c inversify.ContainerBinder) error {
			c.Unbind(auth.UserRepositoryServiceSymbol)
			c.Unbind(auth.JwtRepositoryServiceSymbol)
			c.Unbind(auth.PasswordServiceSymbol)
			c.Unbind(auth.JwtServiceSymbol)

			c.Unbind(auth.DomainEventsSymbol)

			c.Unbind(auth.UserAppSymbol)
			c.Unbind(auth.JwtAppSymbol)

			return nil
		})
}
