// repositories/milk_repo.go - Repository for milk collection
package repository

import (
	"dairy-management-backend/config"
	"dairy-management-backend/entities"

	"gorm.io/gorm"
)

type MilkRepository struct {
	DB *gorm.DB
}

func NewMilkRepository() *MilkRepository {
	return &MilkRepository{DB: config.DB}
}

func (r *MilkRepository) CreateMilkRecord(milk *entities.MilkCollection) error {
	return r.DB.Create(milk).Error
}

func (r *MilkRepository) GetAllMilkRecords() ([]entities.MilkCollection, error) {
	var records []entities.MilkCollection
	err := r.DB.Find(&records).Error
	return records, err
}

func (r *MilkRepository) GetMilkRecordByID(id uint) (*entities.MilkCollection, error) {
	var record entities.MilkCollection
	err := r.DB.First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *MilkRepository) UpdateMilkRecord(milk *entities.MilkCollection) error {
	return r.DB.Save(milk).Error
}

func (r *MilkRepository) DeleteMilkRecord(id uint) error {
	return r.DB.Delete(&entities.MilkCollection{}, id).Error
}
