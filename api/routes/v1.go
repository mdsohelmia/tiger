package routes

import (
	"github.com/MdSohelMia/tiger/api/controllers"
	"github.com/MdSohelMia/tiger/api/controllers/library"
	"github.com/MdSohelMia/tiger/api/controllers/video"
	"github.com/MdSohelMia/tiger/api/middleware"
	"github.com/MdSohelMia/tiger/core"
	"go.uber.org/zap"
)

type ApiRoutes struct {
	app               *core.App
	logger            *zap.SugaredLogger
	authMiddleware    *middleware.AuthMiddleware
	indexController   *controllers.IndexController
	videoController   video.VideoController
	libraryController *library.LibraryController
}

func (api ApiRoutes) Setup() {

	health := api.app.Kernel
	health.GET("/", api.indexController.Index)
	health.GET("/health-check", api.indexController.HealthCheck)
	// V1 Group
	v1 := health.Group("v1").Use(api.authMiddleware.Authenticate())
	//Library Management Routes
	v1.GET("/libraries", api.libraryController.Index)
	v1.POST("/libraries", api.libraryController.Store)
	v1.GET("/libraries/:id", api.libraryController.Show)
	v1.PUT("/libraries/:id", api.libraryController.Show)

	//Collection Management Routes
	v1.GET("/collections", api.libraryController.Index)
	v1.POST("/collections", api.libraryController.Store)
	v1.GET("/collections/:id", api.libraryController.Show)
	v1.PUT("/collections/:id", api.libraryController.Show)

	//Video Management Routes
	v1.GET("/videos", api.videoController.Index)
	v1.POST("/videos", api.videoController.Create)
	v1.GET("/videos/:id", api.videoController.Show)
	v1.PUT("/videos/:id", api.videoController.Show)

	//Video Player
	v1.GET("/players", api.videoController.Index)
	v1.POST("/players", api.videoController.Create)
	v1.GET("/players/:id", api.videoController.Show)
	v1.PUT("/players/:id", api.videoController.Show)

	//Settings
	v1.GET("/settings", api.videoController.Index)
	v1.POST("/settings", api.videoController.Create)
	v1.GET("/settings/:id", api.videoController.Show)
	v1.PUT("/settings/:id", api.videoController.Show)

}

func NewApiRoutes(app *core.App,
	logger *zap.SugaredLogger,
	//Middler
	authMiddleware *middleware.AuthMiddleware,
	//Controller
	indexController *controllers.IndexController,
	videoController video.VideoController,
	libraryController *library.LibraryController,
) ApiRoutes {
	return ApiRoutes{
		app:               app,
		logger:            logger,
		authMiddleware:    authMiddleware,
		videoController:   videoController,
		libraryController: libraryController,
		indexController:   indexController,
	}
}
