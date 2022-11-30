package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gotipath.com/gostream/core"
)

type AuthMiddleware struct {
	app    *core.App
	logger *zap.SugaredLogger
}

func NewAuthMiddleware(app *core.App, logger *zap.SugaredLogger) *AuthMiddleware {
	return &AuthMiddleware{
		app:    app,
		logger: logger,
	}
}

func (m *AuthMiddleware) Setup() {
	m.logger.Info("setting up database transaction middleware")
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		m.logger.Info("Hello Auth middleware")

		ctx.Next()
	}
}
