package repository

import (
	"context"
	"fmt"

	"github.com/MdSohelMia/tiger/core"
	"github.com/MdSohelMia/tiger/src/models"
	"github.com/gin-gonic/gin"
)

type InterfaceLibraryRepository interface {
	All() any
	Create() interface{}
	Destroy() interface{}
	Update() interface{}
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

func (r *LibraryRepository) Create(data models.Library) {

	_, err := r.db.Connection.NewInsert().Model(&data).Exec(r.ctx)

	if err != nil {
		fmt.Printf("error :%v", err.Error())
	}
}

func (r *LibraryRepository) All() ([]models.Library, error) {
	var libraries []models.Library

	err := r.db.Connection.NewSelect().
		Model(&r.model).
		Order("created_at DESC").
		Scan(r.ctx, &libraries)

	if err != nil {
		return []models.Library{}, err
	}

	return libraries, nil
}

func (r *LibraryRepository) Update() {

}

func (r *LibraryRepository) Destroy() {

}

func (r *LibraryRepository) Paginate(ctx gin.Context) {

}

func (r *LibraryRepository) Find(id string) (models.Library, error) {

	var library models.Library

	err := r.db.Connection.NewSelect().Model(&library).Where("id = ?", id).Scan(r.ctx, &library)

	if err != nil {
		return models.Library{}, err
	}
	return library, nil
}
