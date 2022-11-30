package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gotipath.com/gostream/config"
)

type App struct {
	Kernel *gin.Engine
	logger *zap.SugaredLogger
	config *config.Config
}

func NewApp(logger *zap.SugaredLogger, config *config.Config) *App {
	kernel := gin.New()
	// Default Logger & Recovery
	kernel.Use(gin.Logger(), gin.Recovery())
	// App Kernel
	return &App{
		Kernel: kernel,
		logger: logger,
		config: config,
	}
}

// Run Application
func (app *App) Run() {

	app.Kernel.Run(fmt.Sprintf(":%s", app.config.App.Port))
}
