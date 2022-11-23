package middleware

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewMiddlewares),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middleware contains multiple middleware
type Middlewares []IMiddleware

// NewMiddleware creates new middleware
// Register the middleware that should be applied directly (globally)

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}

func NewMiddlewares() Middlewares {
	return Middlewares{}
}
