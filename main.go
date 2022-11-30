package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"gotipath.com/gostream/api/controllers"
	"gotipath.com/gostream/api/middleware"
	"gotipath.com/gostream/api/routes"
	"gotipath.com/gostream/config"
	"gotipath.com/gostream/core"
	"gotipath.com/gostream/src/repository"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			core.Options,
			core.NewApp,
			core.NewDatabase,
		),
		// Middleware provider
		middleware.Module,
		//Routes Provider
		routes.Module,
		//Controller Providers
		controllers.Module,
		//Repository
		repository.Module,

		fx.Invoke(Register),
	).Run()
}

func Register(
	lfc fx.Lifecycle,
	logger *zap.SugaredLogger,
	app *core.App,
	route routes.Routes,
	middleware middleware.Middlewares,
	config *config.Config,
) {
	logger.Info("Hook method run!")
	fmt.Println(config)
	lfc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				//Setup All Routes.
				middleware.Setup()
				route.Setup()
				go app.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)

}
