package repository

import (
	"gotipath.com/gostream/core"
	"gotipath.com/gostream/src/models"
)

type CollectionRepository struct {
	db    *core.Database
	model models.Collection
}

func NewCollectionRepository(db *core.Database) *CollectionRepository {
	return &CollectionRepository{
		db:    db,
		model: models.Collection{},
	}
}

func (collectionRepo *CollectionRepository) Create(data models.Collection) (models.Collection, error) {

	collectionRepo.db.Connection.AutoMigrate(&collectionRepo.model)

	if err := collectionRepo.db.Connection.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (repo *CollectionRepository) List() []models.Collection {
	var collections []models.Collection
	repo.db.Connection.Model(&repo.model).Where("parent_id = ?", "").Limit(10).Find(&collections)
	return collections
}
