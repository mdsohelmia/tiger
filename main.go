package main

import (
	"context"
	"fmt"

	"github.com/MdSohelMia/tiger/api/controllers"
	"github.com/MdSohelMia/tiger/api/middleware"
	"github.com/MdSohelMia/tiger/api/routes"
	"github.com/MdSohelMia/tiger/config"
	"github.com/MdSohelMia/tiger/core"
	"github.com/MdSohelMia/tiger/src/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
