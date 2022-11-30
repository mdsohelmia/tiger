package collection

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gotipath.com/gostream/constants"
	"gotipath.com/gostream/core"
	"gotipath.com/gostream/src/models"
	"gotipath.com/gostream/src/repository"
)

type CollectionController struct {
	collectionRepo *repository.CollectionRepository
}

func NewCollectionController(collectionRepo *repository.CollectionRepository) *CollectionController {
	return &CollectionController{
		collectionRepo: collectionRepo,
	}
}

func (ctrl *CollectionController) Index(c *gin.Context) {

	collection := ctrl.collectionRepo.List()

	c.JSON(http.StatusOK, gin.H{
		"message": "List",
		"data":    collection,
	})
}

func (ctrl *CollectionController) Store(c *gin.Context) {
	var request CollectionCreateRequest

	if err := c.BindJSON(&request); err != nil {
		core.ResponseWithError(c, http.StatusBadRequest, err.Error(), constants.JSON_MALFORMED)
		return
	}

	//Validation Handle
	if err := request.Validate(); err != nil {
		errors := core.ValidationErrors(err.(validation.Errors))

		core.ResponseWithError(c, http.StatusUnprocessableEntity, "validation errors", errors)
		return
	}

	data, err := ctrl.collectionRepo.Create(models.Collection{
		Name:      request.Name,
		LibraryId: request.LibraryId,
		ParentId:  request.ParentId,
	})

	if err != nil {
		core.ResponseWithError(c, http.StatusBadRequest, err.Error(), constants.JSON_MALFORMED)
		return
	}

	core.ResponseWithSuccess(c, "collections", data)
}
