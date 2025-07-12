package record

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        *string         `json:"id,omitempty"`
	CreatedAt *time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time      `gorm:"autoCreateTime" json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	var newID = uuid.New().String()
	b.ID = &newID

	return nil
}
