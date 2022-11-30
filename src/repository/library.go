package repository

import (
	"context"

	"github.com/gin-gonic/gin"
	"gotipath.com/gostream/core"
	"gotipath.com/gostream/src/models"
)

type InterfaceLibraryRepository interface {
	All() ([]models.Library, error)
	Create(data models.Library) (interface{}, error)
	Destroy(id string)
	Update()
	Paginate(ctx gin.Context) interface{}
}

type LibraryRepository struct {
	db    *core.Database
	ctx   context.Context
	model models.Library
}

func NewLibraryRepository(db *core.Database) *LibraryRepository {
	return &LibraryRepository{
		db:    db,
		ctx:   context.Background(),
		model: models.Library{},
	}
}

func (r *LibraryRepository) Create(data models.Library) (interface{}, error) {

	if err := r.db.Connection.Model(&r.model).Create(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *LibraryRepository) All() ([]models.Library, error) {
	var libraries []models.Library

	r.db.Connection.Order("created_at desc").Find(&libraries)

	return libraries, nil
}

func (r *LibraryRepository) Update() {

}

func (r *LibraryRepository) Destroy(id string) (bool, error) {

	if err := r.db.Connection.Where("id", id).Delete(&models.Library{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *LibraryRepository) Paginate(ctx gin.Context) {

}

func (r *LibraryRepository) Find(id string) (models.Library, error) {
	var library models.Library
	r.db.Connection.First(&library, "id = ?", id)
	return library, nil
}
