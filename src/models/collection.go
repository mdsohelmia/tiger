package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Collection
type Collection struct {
	ID            string    `gorm:"type:UUID;primaryKey;default:generateUUIDv4()" json:"id,omitempty"`
	CustomerId    string    `json:"customer_id,omitempty"`
	LibraryId     string    `json:"library_id"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	DeletedAt     string    `json:"deleted_at,omitempty"`
	Name          string    `json:"name,omitempty"`
	ParentId      string    `json:"parent_id,omitempty"`
	VideoCount    uint64    `json:"video_count" gorm:"-"`
	ChildrenCount uint16    `json:"children_count" gorm:"-"`
}

func (lib *Collection) BeforeCreate(tx *gorm.DB) (err error) {
	lib.ID = uuid.New().String()

	return
}
