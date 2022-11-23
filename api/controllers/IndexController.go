package controllers

import (
	"net/http"

	"github.com/MdSohelMia/tiger/core"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	app *core.App
}

func (ctrl *IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Welcome Gotipath Cloud & Stream service.",
	})
}

func (ctrl *IndexController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All Service available.",
	})
}

func NewIndexController(app *core.App) *IndexController {
	return &IndexController{
		app: app,
	}
}
