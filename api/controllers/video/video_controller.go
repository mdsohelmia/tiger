package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gotipath.com/gostream/core"
	"gotipath.com/gostream/src/repository"
)

type VideoController struct {
	app  *core.App
	repo *repository.LibraryRepository
}

func NewVideoController(
	app *core.App,
	repo *repository.LibraryRepository,
) VideoController {
	return VideoController{
		app:  app,
		repo: repo,
	}
}

func (ctrl *VideoController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
}

func (ctrl *VideoController) Create(c *gin.Context) {

}

func (ctrl *VideoController) Update(c *gin.Context) {

}

func (ctrl *VideoController) Show(c *gin.Context) {

}
func (ctrl *VideoController) Destroy(c *gin.Context) {

}
