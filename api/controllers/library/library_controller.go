package library

import (
	"fmt"
	"net/http"

	"github.com/MdSohelMia/tiger/core"
	"github.com/MdSohelMia/tiger/src/models"
	"github.com/MdSohelMia/tiger/src/repository"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

// library_controller
type LibraryController struct {
	app  *core.App
	repo *repository.LibraryRepository
}

func NewLibraryController(app *core.App, repo *repository.LibraryRepository) *LibraryController {
	return &LibraryController{
		app:  app,
		repo: repo,
	}
}

func (ctrl *LibraryController) Index(c *gin.Context) {
	data, err := ctrl.repo.All()

	if err != nil {
		fmt.Println(err)
	}

	core.ResponseWithSuccess(c, "Library List", data)

}

func (ctrl *LibraryController) Store(c *gin.Context) {

	var request CreateLibraryRequest

	if err := c.BindJSON(&request); err != nil {
		core.ResponseWithError(c, http.StatusBadRequest, err.Error(), "error")
		return
	}

	//Validation Handle
	if err := request.Validate(); err != nil {
		errors := core.ValidationErrors(err.(validation.Errors))

		core.ResponseWithError(c, http.StatusUnprocessableEntity, "validation errors", errors)
		return
	}

	ctrl.repo.Create(models.Library{
		ID:         uuid.New(),
		Title:      request.Title,
		CustomerId: request.CustomerId,
	})

	core.ResponseWithSuccess(c, "Library List", request)
}

func (ctrl *LibraryController) Show(c *gin.Context) {
	libId := c.Param("id")
	data, err := ctrl.repo.Find(libId)

	if err != nil {
		core.ResponseWithError(c, http.StatusNotFound, fmt.Sprintf("This `%s` not found library", libId), []string{})
		return
	}

	core.ResponseWithSuccess(c, "Library details", data)
}

func (ctrl *LibraryController) Destroy(c *gin.Context) {
	core.ResponseWithSuccess(c, "Library List", []string{})
}

func (ctrl *LibraryController) Update(c *gin.Context) {
	core.ResponseWithSuccess(c, "Library List", []string{})
}
