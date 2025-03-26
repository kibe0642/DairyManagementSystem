package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cow struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	TagID     string         `gorm:"unique;not null" json:"tag_id"`
	Breed     string         `json:"breed"`
	Age       float64        `json:"age"`
	AgeUnit   string         `json:"age_unit"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
