package repository

import (
	"dairy-management-backend/entities"

	"gorm.io/gorm"
)

type CowRepository struct {
	DB *gorm.DB
}

// ✅ Now accepts *gorm.DB as an argument
func NewCowRepository(db *gorm.DB) *CowRepository {
	return &CowRepository{DB: db}
}

func (r *CowRepository) CreateCow(cow *entities.Cow) error {
	return r.DB.Omit("ID").Create(cow).Error
}
func (r *CowRepository) GetAllCows() ([]entities.Cow, error) {
	var cows []entities.Cow
	err := r.DB.Find(&cows).Error
	return cows, err
}

func (repo *CowRepository) GetCowByTagID(tagID string) (*entities.Cow, error) {
	var cow entities.Cow
	if err := repo.DB.Where("tag_id = ?", tagID).First(&cow).Error; err != nil {
		return nil, err
	}
	return &cow, nil
}

func (r *CowRepository) UpdateCow(cow *entities.Cow) error {
	return r.DB.Save(cow).Error
}

func (r *CowRepository) DeleteCow(id uint) error {
	return r.DB.Delete(&entities.Cow{}, id).Error
}
